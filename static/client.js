window.onload = function () {

    const subButton = document.querySelectorAll("button")[0]
    const formInput = document.querySelectorAll("input")[0]
    console.log(subButton)
    subButton.onclick = async function (e) {
        e.preventDefault()
        let val = formInput.value

        let body = { val }
        let response = await fetch("http://localhost:3000//snip", {
            method: "POST",
            body: JSON.stringify(body),
            headers: {
                "Content-Type": "application/json"
            }
        })
        console.log("SENT: " + val)
        console.log(await response.json())

    }
}