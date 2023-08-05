<script>
import { Router, Link, Route } from "svelte-routing";
import Home from "./pages/Home.svelte";
import Register from "./pages/Register.svelte";
import Login from "./pages/Login.svelte";
import Admin from "./pages/Admin.svelte";
import CreateSurvey from "./pages/survey/Create.svelte";
import Survey from "./pages/survey/Survey.svelte";

export let url = "/";
import Cookies from "js-cookie";
let User = false;
let user = Cookies.get("Token");
async function GetUser(){
url = window.location.pathname;
if(user){
  const response = await fetch(`http://localhost:3000/find/user/${user}`, {
      headers: {
        "Content-Type": "application/json",
      }
    });

    if (response.ok) {
      const data = await response.json();
      User = data.data;
      console.log(data)
    } else if (response.status === 404) {
      console.log("Kullanıcı bulunamadı.");
    }
  } 
}
GetUser();

</script>


<Router {url}>
<nav>
  <Link to="/">Anasayfa</Link>
  {#if User}
    {User.Username}
    <Link to="/create/survey">Anket oluştur</Link>
    
    <Route path="/create/survey" component={CreateSurvey} />
    <Route path="/survey/:token" let:params >
        <Survey token={params.token} user={User}/>
    </Route>
    {#if User.Perm == "Admin"}
      <Link to="/admin/dashboard">Admin Dashboard</Link>
      <Route path="/admin/dashboard" component={Admin} />
    {/if}
  {:else}
  <Link to="/register">Kayıt ol</Link>
  <Link to="/login">Giriş yap</Link>
  {/if}
</nav>
<Route path="/" component={Home} />
<Route path="/register" component={Register} />
<Route path="/login" component={Login} />
</Router>