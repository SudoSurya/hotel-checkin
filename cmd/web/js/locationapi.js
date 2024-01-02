async function getCountries() {
    try {
        const response = await fetch("https://api.countrystatecity.in/v1/countries", {
            method: 'GET',
            headers: {
                'X-CSCAPI-KEY': 'd3ZwTnBXVERmSXlMNzd3dlc5SXBRS29uR2plcEVVeGxwZ3lWQzhvQg==', // Replace with your API key
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch countries');
        }

        const data = await response.json();
        const countriesDropdown = document.getElementById("country");

        // Clear existing options
        countriesDropdown.innerHTML = "";

        // Add a default option
        const defaultOption = document.createElement("option");
        defaultOption.value = ""; // Set an appropriate default value
        defaultOption.text = "Select Country"; // Set an appropriate default text
        countriesDropdown.appendChild(defaultOption);

        // Add options for each country
        data.forEach(country => {
            const option = document.createElement("option");
            option.value = country.iso2; // Set the value to the country ID or relevant identifier
            option.text = country.name; // Set the text to the country name
            countriesDropdown.appendChild(option);
        });

        // Trigger the updateStates function to populate the states dropdown based on the selected country
        // updateStates();
    } catch (error) {
        console.error('Error fetching countries:', error);
    }
}

async function getStates(country) {
    try {
        const response = await fetch(`https://api.countrystatecity.in/v1/countries/${country}/states`, {
            method: 'GET',
            headers: {
                'X-CSCAPI-KEY': 'd3ZwTnBXVERmSXlMNzd3dlc5SXBRS29uR2plcEVVeGxwZ3lWQzhvQg==', // Replace with your API key
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch states');
        }

        const data = await response.json();
        const statesDropdown = document.getElementById("state");

        // Clear existing options
        statesDropdown.innerHTML = "";

        // Add options for each state
        data.forEach(state => {
            const option = document.createElement("option");
            option.value = state.iso2; // Set the value to the state ID or relevant identifier
            option.text = state.name; // Set the text to the state name
            statesDropdown.appendChild(option);
        });

        // Trigger the updateCities function to populate the cities dropdown based on the selected state
        // updateCities();
    } catch (error) {
        console.error('Error fetching states:', error);
    }
}

async function getCities(state, country) {
    try {
        const response = await fetch(`https://api.countrystatecity.in/v1/countries/${country}/states/${state}/cities`, {
            method: 'GET',
            headers: {
                'X-CSCAPI-KEY': 'd3ZwTnBXVERmSXlMNzd3dlc5SXBRS29uR2plcEVVeGxwZ3lWQzhvQg==', // Replace with your API key
            },
        });

        console.log("in get cities");

        if (!response.ok) {
            throw new Error(`Failed to fetch cities. Status: ${response.status} ${response.statusText}`);
        }

        const data = await response.json();
        console.log("Cities data:", data);

        const citiesDropdown = document.getElementById("city");

        // Clear existing options
        citiesDropdown.innerHTML = "";

        // Add options for each city
        data.forEach(city => {
            const option = document.createElement("option");
            option.value = city.id; // Set the value to the city ID or relevant identifier
            option.text = city.name; // Set the text to the city name
            citiesDropdown.appendChild(option);
        });
    } catch (error) {
        console.error('Error fetching cities:', error.message);
    }
}
function updateStates() {
    var selectedCountry = document.getElementById("country").value;
    console.log(selectedCountry);
    getStates(selectedCountry);
}

function updateCities() {
    var selectedState = document.getElementById("state").value;
    var selectedCountry = document.getElementById("country").value;
    console.log(selectedState);
    console.log(selectedCountry);
    getCities(selectedState, selectedCountry);
}
getCountries();

document.addEventListener("DOMContentLoaded", function() {
    var form = document.getElementById("hotelrefform");
    form.reset();
});
