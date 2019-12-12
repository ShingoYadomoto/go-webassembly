var img = new Image();
img.src = "./image/gopher.png";

function test(event) {
    var canvas = document.getElementById('sample');
    var context = canvas.getContext('2d');
    context.drawImage(img, event.pageX,event.pageY);
    invert(context, canvas);
}
function invert(ctx, canvas) {
    var imageData = ctx.getImageData(0, 0, canvas.width, canvas.height);
    var data = imageData.data;

    console.log(data);
    for(var i = 0; i < data.length; i += 4) {
        data[i] = 255 - data[i];  // red
        data[i + 1] = 255 - data[i + 1];   // green
        data[i + 2] = 255 - data[i + 2];  // blue
    }
    console.log(data);
    console.log('done');
    ctx.putImageData(imageData, 0, 0);
}