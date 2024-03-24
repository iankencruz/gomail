// Replace div with whats in message textarea
window.onload = function() {
  var htmlbox = document.getElementById("message");

  htmlbox.onkeyup = htmlbox.onkeypress = function() {
    document.getElementById("previewbox").innerHTML = this.value;
  }
}
