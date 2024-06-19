const hashbtn = document.getElementById("execute");
const autoUpdate = document.getElementById("auto-update");

function hashValue(algo, val) {
    return crypto.subtle
        .digest(algo, new TextEncoder('utf-8').encode(val))
        .then(h => {
            let hexes = [],
                view = new DataView(h);
            for (let i = 0; i < view.byteLength; i += 4)
                hexes.push(('00000000' + view.getUint32(i).toString(16)).slice(-8));
            return hexes.join('');
        });
}

function calcHash() {
    const input = document.getElementById('input');
    const algorightm = document.getElementById('algo-type');

    hashValue(algorightm.value, input.value)
        .then((hash) => {
            const output = document.getElementById('output');
            output.value = hash;
        });
}

hashbtn.onclick = () => {
    calcHash()
}

input.oninput = () => {
    if (autoUpdate.checked) {
        calcHash();
    }
}