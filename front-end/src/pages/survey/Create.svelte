<script>
    let title = "";
    let description = "";
    let inputs = [{id:0,name:""}];
    let option = [];

    function addOptions(){
        inputs.push({
            id:inputs.length, 
            name:""
        })
        inputs = inputs;
    }

    function deleteOptions(id){
        inputs = inputs.filter(item => item.id !== id);
    }

    async function send(){
        for(let i=0;i<inputs.length;i++){
            option.push(inputs[i].name)
        }

        let options = option.join(",");
    
        const requestOptions = {
            method: "POST",
            mode:"cors",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: 'include',
            body: JSON.stringify({
                title,
                description,
                options,
            }),
        };

    try {
        const response = await fetch("http://localhost:3000/api/create/survey", requestOptions);
        const data = await response.json();
        console.log(data.message); 
        } catch (error) {
            console.error("An error occurred:", error);
        }
    }


</script>
<main>
    <input bind:value={title} required /><br>
    <input bind:value={description} required /><br>
    {#each inputs as input}
        <input type="text" bind:value={input.name}>
        <button on:click={deleteOptions(input.id)}>Seçeneği sil</button><br>
    {/each}
    <button on:click={addOptions}>Seçenek ekle</button><br>
    <button on:click={send}>gönder</button>
</main>