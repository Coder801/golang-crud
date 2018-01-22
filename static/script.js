document.addEventListener("DOMContentLoaded", () => {

  window.renderData = function(data) {
    const list = document.querySelector('.employees');
    const html = data.map(item => `<tr>
        <td>${item.id}</td>
        <td>${item.username}</td>
        <td>${item.position}</td>
      </tr>`).join('');

    list.innerHTML = html;
  }

})
