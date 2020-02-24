const axios = require("axios");
const docuSignURL = process.env['DS_GATEWAY_URL'];
console.log(`DocuSign URL = ${docuSignURL}`);

module.exports = {
    main: async function (event, context) {
        try {
            console.log("ce-type: " + event.extensions.request.headers['ce-type']);
            console.log(event.data);
            const envelopeId = event.data.EnvelopeId;

            const response = await axios({
                method: 'get',
                url: `${docuSignURL}/envelopes/${envelopeId}`
            });
            console.log(response.headers);
            console.log(response.data);
        } catch (error) {
            console.log("Error:");
            console.log(error);
            event.extensions.response.status(500).send("Error");
        }
    }
}