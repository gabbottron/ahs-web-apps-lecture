// api.js
document.addEventListener("DOMContentLoaded", () => {
    const itemList = document.getElementById("item-list");

    // Function to fetch and display items
    function fetchItems() {
        fetch("http://localhost:8087/items")
            .then((response) => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then((data) => {
                itemList.innerHTML = ""; // Clear existing list
                data.forEach((item) => {
                    const listItem = document.createElement("li");
                    listItem.innerHTML = `<a href="details.html?item_id=${item.id}">${item.name} - Quantity: ${item.quantity}</a>`;
                    itemList.appendChild(listItem);
                });
            })
            .catch((error) => {
                console.error("Error fetching items:", error);
            });
    }

    // Call the fetchItems function to populate the list
    fetchItems();
});
