const axios = require('axios');

async function get (id) {

    try {
        const response = await axios.get("http://localhost:5008/api/Joiner?id=" + id)
        return response.data
    } catch (err) {
        console.error(err)
        return null;
    }    
}

module.exports = {
    get: get
  }