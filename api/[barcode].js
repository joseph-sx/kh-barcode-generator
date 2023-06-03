const { DOMImplementation, XMLSerializer } = require('xmldom');
const  JsBarcode = require('jsbarcode');

module.exports = (req, res) => {
  const { barcode }  = req.query;
  console.log(req.headers);
  if(new String(barcode).length <= 25){
    const xmlSerializer = new XMLSerializer();
    const document = new DOMImplementation().createDocument('http://www.w3.org/1999/xhtml', 'html', null);
    const svgNode = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    JsBarcode(svgNode, barcode, {
      xmlDocument: document,
      width:1,
      displayValue: false,
      fontSize: 15,
      height: 25,
      margin:0
    });
    const svgText = xmlSerializer.serializeToString(svgNode);
    res.setHeader('content-type', 'image/svg+xml');
    res.send(svgText);
  }else{
    res.status(500);
    res.send('ERROR: Invalid, Input must be <= 25');
  }
};
