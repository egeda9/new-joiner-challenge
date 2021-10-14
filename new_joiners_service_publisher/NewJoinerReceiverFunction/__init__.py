import logging
import azure.functions as func
from itertools import *
import spacy
nlp = spacy.load("en_core_web_trf")
import PyPDF2
import pika
import docx2txt
import json
from json import JSONEncoder
import yaml

ALLOWED_EXTENSIONS = {'pdf', 'docx'}

class Ner: 
    def __init__(self, tag, value): 
        self.tag = tag 
        self.value = value

class NerEncoder(JSONEncoder):
        def default(self, o):
            return o.__dict__
        
def get_config():
    
    with open("config.yml", "r") as config:
        content = yaml.load(config)
        
    return content

def get_ner(text):

    doc = nlp(text)
    entities = [] 

    for chunk in doc.noun_chunks:
        entities.append(Ner('Noun_Phrase', chunk.text.strip()))

    for token in doc:
        if token.pos_ == "VERB":
            entities.append(Ner('Verb', token.lemma_.strip()))

    for word in doc.ents:
        entities.append(Ner(word.label_, word.text.strip()))
        
    return entities

def read_pdf(file):
    
    text = ''    

    reader = PyPDF2.PdfFileReader(file)
    for page in reader.pages:
        text += page.extractText()
            
    return text.strip()

def send_message(message):
    
    config = get_config()
    connection = pika.BlockingConnection(pika.ConnectionParameters(config["server"]["broker"]))
    channel = connection.channel()

    channel.queue_declare(queue='joiner')

    channel.basic_publish(exchange='',
                          routing_key='joiner',
                          body=message)

    connection.close()

def read_docx(file):
    text = docx2txt.process(file)          
    return text.strip()

def allowed_file(filename):
    return '.' in filename and \
           filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS

def main(req: func.HttpRequest) -> func.HttpResponse:
    logging.info('Python HTTP trigger function processed a request.')
 
    try:
        if req.method == 'POST':
            if 'file' not in req.files:
                logging.info('No file part')                
                return func.HttpResponse("No file part", status_code=400)

            file = req.files['file']

            if file.filename == '':
                logging.info('No selected file')                
                return func.HttpResponse("No selected file", status_code=400)

            if file and allowed_file(file.filename):
                
                if file.content_type == 'application/vnd.openxmlformats-officedocument.wordprocessingml.document':
                    text = read_docx(file)
                elif file.content_type == 'application/pdf':
                    text = read_pdf(file)
                else:
                    raise Exception('Invalid file type ' + file.content_type)
                
                if(not text):
                    return func.HttpResponse("Message was not sent. Empty file content", status_code=204)
                else:
                    entities = get_ner(text)

                    result = dict()
                    for key, g in groupby(sorted(entities, key=lambda x: x.tag), lambda x: x.tag):
                        seen = set()
                        l = []
                        for ent in list(g):
                            if ent.value not in seen:
                                seen.add(ent.value)
                                l.append(ent.value)
                        result[key] = l

                    jsonMessage = json.dumps(result)
                    send_message(jsonMessage)                                
                             
                return func.HttpResponse("OK", status_code=200)

            else:      
                logging.info('Invalid request action')                
                return func.HttpResponse("Invalid request action", status_code=404)          

    except Exception as e:
        raise Exception(str(e))