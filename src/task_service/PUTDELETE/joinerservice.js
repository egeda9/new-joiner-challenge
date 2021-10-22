const axios = require('axios');

const config = {
    joinerBaseUrl: process.env.JOINER_BASEURL
  }

async function get (id) {

    try {
        const response = await axios.get(config.joinerBaseUrl +  "/api/Joiner?id=" + id)
        return response.data
    } catch (err) {
        console.error(err)
        return null;
    }    
}

module.exports = {
    get: get
  }