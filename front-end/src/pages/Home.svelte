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
    <a href="/survey/{survey.Token}" class="py-5 px-5 rounded-3xl bg-gray-600 hover:outline-none border-4 border-blue-600">{survey.Title}</a><br>
{/each}
</main> 
