# VehicleTracker

<p>Repositório do código-base da app de <strong>Rastreamento Veicular Web</strong>.</p>

<h3>Techs</h3>
<ul>
<li><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/9a/Visual_Studio_Code_1.35_icon.svg/2048px-Visual_Studio_Code_1.35_icon.svg.png" width="5%" height="5%"> IDE para desenvolvimento (Usarei VsCode utilizando a extensão devcontainer)</li>
<li><img src="https://www.tshirtgeek.com.br/wp-content/uploads/2021/08/com031.jpg" width="5%" heidht="5%"> Banco de Dados MySQL</li>
<li><img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="5%" heidht="5%"> Golang go1.19.8</li>
<li><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a8/NestJS.svg/1200px-NestJS.svg.png" width="5%" heidht="5%"> NestJs para back-end</li>
<li><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--RpUfSAFP--/c_imagga_scale,f_auto,fl_progressive,h_1080,q_auto,w_1080/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/8otweo5ef6kwc26rmxe5.png" width="5%" heidht="5%"> NextJs para front-end</li>
<li><img src="https://i1.wp.com/airisdata.com/wp-content/uploads/2016/01/kafka-logo-600x390.jpg?fit=600%2C390" width="5%" heidht="5%"> Kafka para controle de tópico que será consumido pelo Golang</li>
<li><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/38/Prometheus_software_logo.svg/2066px-Prometheus_software_logo.svg.png" width="5%" heidht="5%"> Prometheus para integrar ao Grafana</li>
<li><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a1/Grafana_logo.svg/2005px-Grafana_logo.svg.png" width="5%" heidht="5%"> Grafana para dash de monitoramento</li>
</ul>

<p>Utilizado api <strong>Places API</strong> da Google para conusmo dos mapas/locais, recomedado que seja criado uma key particular para utilizá-la. (Importante ser inserida a key no .env do Nest e Next como valor da varável <strong>GOOGLE_MAPS_API_KEY</strong>).</p>
<p>Para pagina de admin temos comunicação via websocket para constante atualização da localização do veículo.</p>
<p>Para iniciar os containers está disponibilizado o <strong>run.sh</strong> que irá subir os containers </p>


<h3>Author</h3>

<a href="https://www.linkedin.com/in/renato-luciano-barbosa-fh/">
 <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/u/56805018?v=4" width="100px;" alt=""/>
 <br />
 
 <sub><b>Renato Barbosa</b></sub></a>
 
 
Entre em contato!

[![Linkedin Badge](https://img.shields.io/badge/-Renato-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/renato-luciano-barbosa-fh/)](https://www.linkedin.com/in/renato-luciano-barbosa-fh/)
[![Gmail Badge](https://img.shields.io/badge/-renatoinfbarbosa@gmail.com-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:renatoinfbarbosa@gmail.com)](mailto:renatoinfbarbosa@gmail.com)
<hr>
