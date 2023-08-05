<script>
    let surveys = [];
    let error;
    async function getSurveys(){
        fetch(`http://localhost:3000/surveys`).then(async(data)=>{
            if(data.ok){
                let res = await data.json();
                surveys = res.data;
            } else {
                error = "anket bulunamadı"
            }
        })
    }

    getSurveys()
</script>
<main>
<h2>Anketler</h2>
{#each surveys as survey}
    <a href="/survey/{survey.Token}">{survey.Title}</a><br>
{/each}
</main> 