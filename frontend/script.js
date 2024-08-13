const units = {
    length: ["Millimeter", "Centimeter", "Meter", "Kilometer", "Inch", "Foot", "Yard", "Mile"],
    weight: ["Milligram", "Gram", "Kilogram", "Ounce", "Pound"],
    temperature: ["Celsius", "Fahrenheit", "Kelvin"]
};

const sectionSelect = document.getElementById("section");
const fromUnitSelect = document.getElementById("fromUnit");
const toUnitSelect = document.getElementById("toUnit");
const form = document.getElementById("converterForm");

function updateUnit(section) {
    fromUnitSelect.innerHTML = ""
    toUnitSelect.innerHTML = ""

    const selectedUnits = units[section]

    selectedUnits.forEach(unit => {
        const opt1 = document.createElement("option")
        opt1.value = unit.toLowerCase()
        opt1.textContent = unit
        fromUnitSelect.appendChild(opt1)

        const opt2 = document.createElement("option")
        opt2.value = unit.toLowerCase()
        opt2.textContent = unit
        toUnitSelect.appendChild(opt2)
    });
}

async function fetchConversionResult(section, fromUnit, toUnit, amount) {
    const url = `http://localhost:6868/convert?section=${section}&from=${fromUnit}&to=${toUnit}&amount=${amount}`;

    // Send a GET request
    try {
        let resp = await fetch(url);
        let data = await resp.json();
        console.log(data.result)
        return data.result;
    } catch (error) {
        console.error('There has been a problem with your fetch operation:', error);
    }
}

sectionSelect.addEventListener("change", () => {
    updateUnit(this.value);
});

form.addEventListener("submit", async (e) => {
    e.preventDefault()

    const section = sectionSelect.value;
    const fromUnit = fromUnitSelect.value;
    const toUnit = toUnitSelect.value;
    const amount = parseFloat(document.getElementById("amount").value);

    const result = await fetchConversionResult(section, fromUnit, toUnit, amount)

    alert(`${amount} ${fromUnit} = ${result} ${toUnit}`)
});

updateUnit(sectionSelect.value)
