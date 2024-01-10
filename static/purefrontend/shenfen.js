// 获取元素
const fileSelector = document.getElementById('imgFile');
const canvas = document.getElementById('canvas');
const textarea = document.getElementById('watermarkText');
const watermark = document.getElementById('watermark')
const downloadBtn = document.getElementById('imgSave');
const imgSelector = document.getElementById('imgSelector');
const refreshBtn = document.getElementById('refresh');

let dataURI = null;

imgSelector.onclick = function (e) {
    fileSelector.click();
}

refreshBtn.onclick = () => {
    if (!!!dataURI) {
        console.log("No image is selected.");
        return;
    }
    loadImage(dataURI);
}

// 处理选择文件
fileSelector.onchange = function (e) {
    const file = e.target.files[0];

    // 文件转为data URI
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onloadend = function () {
        dataURI = reader.result;
        loadImage(dataURI);
    }
}

function hexToRGB(hex, alpha) {
    var r = parseInt(hex.slice(1, 3), 16),
        g = parseInt(hex.slice(3, 5), 16),
        b = parseInt(hex.slice(5, 7), 16);

    if (alpha) {
        return "rgba(" + r + ", " + g + ", " + b + ", " + alpha + ")";
    } else {
        return "rgb(" + r + ", " + g + ", " + b + ")";
    }
}

// 加载图像到canvas
function loadImage(dataURI) {
    const img = new Image();
    img.src = dataURI;
    img.onload = function () {

        watermark.width = img.width / 4;
        watermark.height = img.height / 4;
        
        let transparent = document.getElementById('transparent');
        let textColor = document.getElementById('textColor');
        
        let color = hexToRGB(textColor.value, transparent.value);

        var ctx = watermark.getContext("2d");
        //清除小画布
        ctx.clearRect(0, 0, watermark.width, watermark.height);
        ctx.font = '16px serif';
        //文字倾斜角度
        ctx.rotate(-20 * Math.PI / 180);

        ctx.fillStyle = color;
        //第一行文字
        ctx.fillText(textarea.value, -20, watermark.height * 0.8);
        //坐标系还原
        ctx.rotate(20 * Math.PI / 180);

        canvas.width = img.width;
        canvas.height = img.height;
        var ctxr = canvas.getContext("2d");
        //清除整个画布
        ctxr.clearRect(0, 0, canvas.width, canvas.height);
        ctxr.drawImage(img, 0, 0);
        //平铺--重复小块的canvas
        var pat = ctxr.createPattern(watermark, "repeat");
        ctxr.fillStyle = pat;

        ctxr.fillRect(0, 0, canvas.width, canvas.height);
    }
}

// 下载处理
downloadBtn.onclick = function () {
    const url = canvas.toDataURL('image/png');
    const link = document.createElement('a');
    link.download = 'image.png';
    link.href = url;
    link.click();
}