'use client'
import Grid2 from "@mui/material/Unstable_Grid2/Grid2";
import { Button, CardActions } from "@mui/material";
import LocIcon from "@mui/icons-material/PinDrop";
import RoadIcon from "@mui/icons-material/AddRoad";
import MapIcon from "@mui/icons-material/Monitor";


export default function Home() {
  const handleNovaRotaClick = () => {
    window.location.href = '/new-route';
  };
  const handleDriverClick = () => {
    window.location.href = '/driver';
  };
  const handleAdminClick = () => {
    window.location.href = '/admin';
  };

  return (
    <Grid2 sx={{display: "flex", flex: 1, alignItems: "center", justifyContent: "center", height: "100vh"}} xs={4} px={2}>
      <CardActions sx={{ display: "flex", justifyContent: "center" }}>
        <Button type="button" variant="contained" onClick={handleNovaRotaClick}>
          <RoadIcon />
          Cadastrar nova rota
        </Button>
        <Button type="button" variant="contained" onClick={handleDriverClick}>
          <LocIcon />
          Iniciar rota
        </Button>
        <Button type="button" variant="contained" onClick={handleAdminClick}>
          <MapIcon />
          Monitoramento de rota
        </Button>
      </CardActions>
    </Grid2>
  )
}

