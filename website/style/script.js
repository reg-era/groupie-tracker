function initMap() {
    var map = new google.maps.Map(document.getElementById('map'), {
      center: { lat: 23.98955292905143, lng: 9.679765 },
      zoom: 2,
    });
  
    var loc = document.getElementById('loc')
    var locstring = loc.textContent || loc.innerText;
  
    var locstring = locstring.replace(/'/g, "'");
    var locations = JSON.parse(locstring)
  
    locations.forEach(function (location) {
      new google.maps.Marker({
        position: { lat: location.lat, lng: location.lng },
        map: map,
        title: location.title
      });
    });
  }

  // example of data that we need to give
  //  [
  //   {lat: -37.8136, lng: 144.9631, title: 'Victoria, Australia'},
  //   {lat: -33.8688, lng: 151.2093, title: 'New South Wales, Australia'},
  //   {lat: -27.4698, lng: 153.0251, title: 'Queensland, Australia'},
  //   {lat: -36.8485, lng: 174.7633, title: 'Auckland, New Zealand'},
  //   {lat: -7.7956, lng: 110.3695, title: 'Yogyakarta, Indonesia'},
  //   {lat: 48.1482, lng: 17.1067, title: 'Bratislava, Slovakia'},
  //   {lat: 47.4979, lng: 19.0402, title: 'Budapest, Hungary'},
  //   {lat: 53.9006, lng: 27.5590, title: 'Minsk, Belarus'}
  // ];