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
    <input bind:value={title} required class="py-5 px-5 rounded-3xl bg-gray-600 hover:outline-none border-4 border-blue-600"/><br>
    <input bind:value={description} required class="py-5 px-5 rounded-3xl bg-gray-600 hover:outline-none border-4 border-blue-600"/><br>
    {#each inputs as input}
        <input type="text" bind:value={input.name} class="py-5 px-5 rounded-3xl bg-gray-600 hover:outline-none border-4 border-blue-600">
        <button on:click={deleteOptions(input.id)} class="py-5 px-5 rounded-3xl bg-gray-600 hover:outline-none border-4 border-blue-600">Seçeneği sil</button><br>
    {/each}
    <button on:click={addOptions} class="py-5 px-5 rounded-3xl bg-gray-600 hover:outline-none border-4 border-blue-600">Seçenek ekle</button><br>
    <button on:click={send} class="py-5 px-5 rounded-3xl bg-gray-600 hover:outline-none border-4 border-blue-600">gönder</button>
</main>
