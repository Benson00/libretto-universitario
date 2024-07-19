document.getElementById('viewSelector').addEventListener('change', function() {
    const selectedView = this.value;
    document.querySelectorAll('.chart-container, .table-container').forEach(container => {
        container.style.display = 'none';
    });
    document.getElementById(selectedView).style.display = 'block';
});

fetch('/data')
    .then(response => response.json())
    .then(data => {
        const labels = data.Esami.map(esame => esame.Nome);
        const voti = data.Esami.map(esame => esame.Voto);
        const votiLaurea = data.Esami.map((esame, index) => {
            const subArray = data.Esami.slice(0, index + 1);
            const sommaPesata = subArray.reduce((sum, curr) => sum + curr.Voto * curr.Crediti, 0);
            const creditiTotali = subArray.reduce((sum, curr) => sum + curr.Crediti, 0);
            const mediaPonderata = sommaPesata / creditiTotali;
            return mediaPonderata * 11 / 3;
        });

        // Configurazione del primo grafico (andamento dei voti)
        const ctx = document.getElementById('myChart').getContext('2d');
        new Chart(ctx, {
            type: 'line',
            data: {
                labels: labels,
                datasets: [{
                    label: 'Voti',
                    data: voti,
                    borderColor: 'rgba(75, 192, 192, 1)',
                    borderWidth: 2,
                    backgroundColor: 'rgba(75, 192, 192, 0.2)',
                    fill: true
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                scales: {
                    x: {
                        grid: {
                            display: false
                        }
                    },
                    y: {
                        beginAtZero: false,
                        min: 18,
                        max: 30,
                        grid: {
                            color: 'rgba(0, 0, 0, 0.1)'
                        },
                        ticks: {
                            color: '#34495e'
                        }
                    }
                },
                plugins: {
                    legend: {
                        position: 'top',
                        labels: {
                            color: '#34495e'
                        }
                    }
                }
            }
        });

        // Configurazione del secondo grafico (andamento dei voti di laurea)
        const ctxVoti = document.getElementById('votiChart').getContext('2d');
        new Chart(ctxVoti, {
            type: 'line',
            data: {
                labels: labels,
                datasets: [{
                    label: 'Voti di Laurea',
                    data: votiLaurea,
                    borderColor: 'rgba(153, 102, 255, 1)',
                    borderWidth: 2,
                    backgroundColor: 'rgba(153, 102, 255, 0.2)',
                    fill: true
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                scales: {
                    x: {
                        grid: {
                            display: false
                        }
                    },
                    y: {
                        beginAtZero: false,
                        min: 66, // Assumendo che il minimo voto di laurea sia 66
                        max: 110, // Assumendo che il massimo voto di laurea sia 110
                        grid: {
                            color: 'rgba(0, 0, 0, 0.1)'
                        },
                        ticks: {
                            color: '#34495e'
                        }
                    }
                },
                plugins: {
                    legend: {
                        position: 'top',
                        labels: {
                            color: '#34495e'
                        }
                    }
                }
            }
        });
    })
    .catch(error => console.error('Error fetching data:', error));