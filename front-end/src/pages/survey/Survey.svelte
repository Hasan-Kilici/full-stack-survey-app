<script>
    export let user;
    export let token;
    let title = "";
    let description = "";

    let options = [];
    let error;
    async function FindSurvey(token){
        fetch(`http://localhost:3000/find/survey/${token}`).then(async(data)=>{
            if(data.ok){
                let res = await data.json();
                title = res.data.Title;
                description = res.data.Description;
            } else {
                error = "anket bulunamadı"
            }
        })
    }

    async function FindOptions(token){
        fetch(`http://localhost:3000/find/options/${token}`).then(async(data)=>{
            if(data.ok){
                let res = await data.json();
                options = res.data;
            } else {
                error = "anket bulunamadı"
            }
        })
    }

    async function VoteOption(surveyToken,optionToken){
        fetch(`http://localhost:3000/give/vote/${surveyToken}/${optionToken}`, {
            method:"POST",
            mode:"cors",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: 'include',
        }).then(async(data)=>{
            if(data.ok){
                window.location.href = `/survey/${surveyToken}`
            }
        })
    }

    async function DeleteSurvey(SurveyToken){
        fetch(`http://localhost:3000/delete/survey/${surveyToken}`, {
            method:"POST",
            mode:"cors",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: 'include',
        }).then(async(data)=>{
            if(data.ok){
                window.location.href = `/survey/${surveyToken}`
            }
        })
    }

    FindSurvey(token)
    FindOptions(token)
</script>
<main>
    {#if user.Perm == "Admin"}
        <button on:click={DeleteSurvey(token)}>Anketi sil</button>
    {/if}
    {#if !error}
        <h3>{title}</h3>
        <p>{description}</p>
        {#each options as option}
            <button on:click={VoteOption(option.SurveyToken,option.Token)}>{option.Title} | {option.Votes}</button><br>
        {/each}
    {:else}
        <h3>404</h3>
        {error}
    {/if}
</main>