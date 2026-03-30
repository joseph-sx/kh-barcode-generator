const { DOMImplementation, XMLSerializer } = require('@xmldom/xmldom');
var JsBarcode = require('jsbarcode');

module.exports = (req, res) => {
  const xmlSerializer = new XMLSerializer();
  const document = new DOMImplementation().createDocument(
    'http://www.w3.org/1999/xhtml',
    'html',
    null
  );

  const svgNode = document.createElementNS('http://www.w3.org/2000/svg', 'svg');

  res.setHeader('content-type', 'image/svg+xml');

  var value = req.query.v;

  if (!value) {
    svgNode.setAttribute('width', '150');
    svgNode.setAttribute('height', '50');

    const textNode = document.createElementNS('http://www.w3.org/2000/svg', 'text');
    textNode.setAttribute('x', '50%');
    textNode.setAttribute('y', '50%');
    textNode.setAttribute('dominant-baseline', 'middle');
    textNode.setAttribute('text-anchor', 'middle');
    textNode.setAttribute('font-size', '12');
    textNode.setAttribute('fill', 'black');
    textNode.setAttribute('font-family', 'Arial, sans-serif');

    textNode.textContent = 'NO VALUE!';

    svgNode.appendChild(textNode);

    const svgText = xmlSerializer.serializeToString(svgNode);
    return res.send(svgText);
  }

  JsBarcode(svgNode, value, {
    xmlDocument: document,
    width: 1,
    displayValue: false,
    fontSize: 15,
    height: 25,
    margin: 0
  });

  const svgText = xmlSerializer.serializeToString(svgNode);
  res.send(svgText);
};