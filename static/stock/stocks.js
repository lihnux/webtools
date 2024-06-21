const addStockBtn = document.getElementById("newStockBtn");
const saveNewStockBtn = document.getElementById("saveNewStockBtn");
const cancelNewStockBtn = document.getElementById("cancelNewStockBtn");

fetchStocks();

addStockBtn.onclick = function (e) {
    const newStockDiv = document.getElementById("newStockDiv");
    newStockDiv.style.display = 'block';
}

saveNewStockBtn.onclick = (e) => {
    const newStockDiv = document.getElementById("newStockDiv");
    newStockDiv.style.display = 'none';

    let stock = {
        name: document.getElementById("name").value,
        symbol: document.getElementById("symbol").value
    }

    fetch("https://localhost/stocks", {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(stock)
    })
    .then(response => {
        if (response.ok) {
            return response;
        } else {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
    })
    .then(data => {
        fetchStocks();
        console.log('Response data:', data);
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

cancelNewStockBtn.onclick = (e) => {
    const newStockDiv = document.getElementById("newStockDiv");
    newStockDiv.style.display = 'none';
}

function fetchStocks() {
    fetch("https://localhost/stocks", {
        method: 'GET'
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
    })
    .then(data => {
        console.log('Response data:', data);
        // 获取列表容器元素
        const listContainer = document.getElementById("list-container");
        listContainer.innerHTML = "";

        // // 生成列表项的 HTML 字符串
        // const listItems = data.map((item) => `<li>${item.Name}</li>`).join("");

        // // 将列表项插入到容器中
        // listContainer.innerHTML = listItems;

        data.forEach(item => {
            const li = document.createElement('li');
            const a = document.createElement('a');
            a.text = item.Name;
            a.dataset.symbol = item.Symbol;
            a.href = "#";
            a.addEventListener('click', handleStockClick);
            li.appendChild(a);

            listContainer.appendChild(li);
          });
      
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

function handleStockClick(e) {
    const symbol = event.currentTarget.dataset.symbol;
    console.log(`点击了列表项, Symbol: ${symbol}`);
}