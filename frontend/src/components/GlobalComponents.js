import "../styles/GlobalStyles.css"

function MobileBar() {
    return (
        <header class="contain__bottom">
            <div class="contain__item">
                <a href="#" class="contain__link" target="_self" rel="next"><img src="images/salvos.svg" alt="Obras Salvos" class="contain__img" />Salvos</a>
            </div>
            <div class="contain__item">
                <a href="#" class="contain__link" target="_self" rel="next"><img src="images/comics.svg" alt="Todas as Obras" class="contain__img" />Biblioteca</a>
            </div>
            <div class="contain__item">
                <a href="index.html" class="contain__link" target="_self" rel="prev"><img src="images/home.svg" alt="Pagina Inical" class="contain__img" />Inicio</a>
            </div>
            <div class="contain__item">
                <a href="historic.html" class="contain__link" target="_self" rel="next"><img src="images/historico.svg" alt="Historico" class="contain__img" />histórico</a>
            </div>
            <div class="contain__item">
                <a href="login.html" class="contain__link" target="_self" rel="next"><img src="images/user.svg" alt="Conta do usuário" class="contain__img" />Login</a>
            </div>
        </header>
    );
}

function Sidebar() {
    return (

        <div class="sidebar">
            <div>
                <div>
                    <a class="sidebar-links" href="index.html"><img src="images/logo-full.png" alt={`Logo de ${process.env.REACT_APP_TITLE_NAME}`} /></a>
                </div>
                <input type="text" name="text" placeholder="Buscar..." />
                <span>
                    <img src="images/user.svg" alt="Fazer Login" /><a class="sidebar-links" href="login.html" >Iniciar Sessão</a></span>
            </div>
            <hr />
            <nav class="navbar">
                <ul>
                    <li><img class="img-icons" src="images/home.svg " alt="inicio" /><a class="sidebar-links" href="index.html">Inicio</a></li>
                    <li><img src="images/comics.svg" alt="obras" /><a class="sidebar-links" href="#">Biblioteca</a></li>
                    <li><img src="images/historico.svg" alt="Historico" /><a class="sidebar-links" href="historic.html">Histórico</a></li>
                    <li><img src="images/discord-icon.svg" alt="" /><a class="sidebar-links" href="#">Discord</a></li>
                    <li><img src="images/salvos.svg" alt="" /><a class="sidebar-links" href="#">Salvos</a></li>
                </ul>
            </nav>
            <div class="credits">
                <footer>
                    <p>
                        Desenvolvido por <a class="github" href="https://github.com/isacsoaresz" target="_blank" rel="external">Isac Soares</a> e <a class="github" href="https://github.com/Rochumback" target="_blank" rel="external">
                            Victor Rochumback</a>
                        <br /><a class="discord" href="https://discord.com/users/451559060618084353">@sxnzin</a> <a class="discord" href="https://discord.com/users/391511301010096128">@rochumback</a>
                    </p>
                </footer>
            </div>
        </div >
    );
}

export { MobileBar, Sidebar };

