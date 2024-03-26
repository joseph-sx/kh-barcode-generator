// const { DOMImplementation, XMLSerializer } = require('xmldom');

// var JsBarcode = require('jsbarcode');
const {
    QRCodeRaw,
    QRCodeSVG,
    QRCodeCanvas,
    QRCodeText,
  } = require('@akamfoad/qrcode');
const { Options } = require('discord.js');

module.exports = (req, res) => {

    var value = req.query.v ? req.query.v : 'NO DATA PROVIDED';
    const qrSVG = new QRCodeSVG(value,{size:30});
        // qrSVG.getDataSize() = 30;
        
    console.log(qrSVG.getDataSize());
    // const dataUrlWithSVGQRCode = qrSVG.toDataUrl();
    const xmlWithQRCode = qrSVG.toString();
    res.setHeader('content-type', 'image/svg+xml');

    res.send(xmlWithQRCode);
};
