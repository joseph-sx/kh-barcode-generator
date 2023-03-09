const { DOMImplementation, XMLSerializer } = require('xmldom');
var JsBarcode = require('jsbarcode');

module.exports = (req, res) => {

  const xmlSerializer = new XMLSerializer();
  const document = new DOMImplementation().createDocument('http://www.w3.org/1999/xhtml', 'html', null);
  const svgNode = document.createElementNS('http://www.w3.org/2000/svg', 'svg');



  res.setHeader('content-type', 'image/svg+xml');

  var value = req.query.v;
  JsBarcode(svgNode, value, {
    xmlDocument: document,
    width:1,
    displayValue: false,
    fontSize: 15,
    height: 25,
    margin:0
  });
  const svgText = xmlSerializer.serializeToString(svgNode);
  res.send(svgText);
};
