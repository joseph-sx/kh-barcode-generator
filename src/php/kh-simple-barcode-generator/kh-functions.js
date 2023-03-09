console.log('KH JS Functions loaded');

function printBarcode(sku) {
    var win = window.open('about:blank', "_new");
    win.document.open();
    win.document.write([
        '<html>',
        '   <head>',
        '   </head>',
        '   <body onload="window.print()" onafterprint="window.close()">',
        '       <img src="https://barcode.khat.es/api/generate?v=' + sku + '"/>',
        '   </body>',
        '</html>'
    ].join(''));
    win.document.close();
}