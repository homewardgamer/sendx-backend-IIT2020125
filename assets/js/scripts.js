document.addEventListener("DOMContentLoaded", function () {
  const form = document.querySelector("form");
  const urlInput = document.querySelector("input[name='url']");
  const feedback = document.createElement("div");

  feedback.style.display = "none";
  feedback.style.marginTop = "20px";
  form.parentNode.insertBefore(feedback, form.nextSibling);

  form.addEventListener("submit", function (event) {
    event.preventDefault();
    feedback.style.display = "none";

    if (!isValidUrl(urlInput.value)) {
      feedback.textContent = "Please enter a valid URL.";
      feedback.style.color = "red";
      feedback.style.display = "block";
      return;
    }

    fetch("/crawl", {
      method: "POST",
      body: new URLSearchParams(new FormData(form)),
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
    })
        .then((response) => {
          if (!response.ok) {
            throw new Error("Network response was not ok");
          }
          return response.text();
        })
        .then((page) => {
          const resultDiv = document.getElementById("result");
          resultDiv.innerHTML = "Crawled Successfully!";
          resultDiv.style.color = "green";
          resultDiv.style.display = "block";
        })
        .catch((error) => {
          feedback.textContent = "There was an unexpected error.";
          feedback.style.color = "red";
          feedback.style.display = "block";
        });
  });

  function isValidUrl(string) {
    try {
      new URL(string);
      return true;
    } catch (_) {
      return false;
    }
  }
});
