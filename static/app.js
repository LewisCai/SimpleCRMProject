document.getElementById('leadForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const name = document.getElementById('name').value;
    const company = document.getElementById('company').value;
    const email = document.getElementById('email').value;
    const phone = document.getElementById('phone').value;

    fetch('/api/v1/lead', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name, company, email, phone })
    })
    .then(response => response.json())
    .then(data => {
        alert('Lead added!');
        fetchLeads();
    })
    .catch(error => console.error('Error adding lead:', error));
});

document.getElementById('deleteSelected').addEventListener('click', function() {
    const selectedLeads = document.querySelectorAll('input[name="leadCheckbox"]:checked');

    if (selectedLeads.length === 0) {
        alert('No leads selected');
        return;
    }

    selectedLeads.forEach(checkbox => {
        const leadId = checkbox.value;
        deleteLead(leadId);
    });
});

function deleteLead(id) {
    console.log('Deleting lead with ID:', id); // Log the ID before the fetch call
    fetch(`/api/v1/lead/${id}`, {
        method: 'DELETE'
    })
    .then(() => {
        fetchLeads(); // Refresh the leads list
    })
    .catch(error => console.error('Error deleting lead:', error));
}


function fetchLeads() {
    fetch('/api/v1/lead')
        .then(response => response.json())
        .then(leads => {
            const leadsList = document.getElementById('leadsList');
            leadsList.innerHTML = ''; // Clear the list before repopulating
            leads.forEach(lead => {
                const listItem = document.createElement('li');

                listItem.innerHTML = `
                    <input type="checkbox" name="leadCheckbox" value="${lead._id}">
                    ${lead.name} - ${lead.company} - ${lead.email} - ${lead.phone}
                `;

                leadsList.appendChild(listItem);
            });
        })
        .catch(error => console.error('Error fetching leads:', error));
}

// Fetch leads on page load
fetchLeads();
