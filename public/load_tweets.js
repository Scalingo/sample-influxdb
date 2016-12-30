document.addEventListener("DOMContentLoaded", function(event) { 
  // Define how to retrieve data for the chart
  var data = {
    url: '/tweets',
    mimeType: 'json',
    xFormat: '%Y-%m-%dT%H:%M:%SZ', // 'xFormat' can be used as custom format of 'x'
    keys: {
      x: 'time',
      value: ['value'],
    },
  };
  // Generate the chart
  var chart = c3.generate({
    bindto: '#chart',
    data: data,
    axis: {
      x: {
        type: 'timeseries',
        tick: {
          format: '%H:%M',
        },
        label: {
          text: "Time",
          position: "center",
        }
      },
      y: {
        min: 0,
        label: {
          text: "Number of tweets",
          position: "outter-middle",
        }
      },
    },
  });

  // Update every second the data
  setInterval(function() {
    chart.load(data);
  }, 1000);
});
