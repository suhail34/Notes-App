<h1>Notes App</h1>
<h3>Pre-Requisites</h3>
<ul>
  <li>
    <p>You need to have <a href="https://go.dev/doc/install">golang version 1.20</a> to run the project locally without docker container</p>
  </li>
  <li>
    <p><a href="https://www.docker.com/get-started/">Docker</a> need to be installed to run it inside a container if you are running it inside the container no need to install golang </p>
  </li>
</ul>
<h3>Run The Program</h3>
<ul>
  <li>
    <p>To run the server locally use <code>make server-local</code></p>
  </li>
  <li>
    <p>To run the server in docker container use <code>make server-container-start</code></p>
  </li>
  <li>
    <p>To stop ther server and delete the container use <code>make server-container-stop</code></p>
  </li>
</ul>
