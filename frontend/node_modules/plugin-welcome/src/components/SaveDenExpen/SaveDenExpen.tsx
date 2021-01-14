import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import { Link as RouterLink } from 'react-router-dom';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Menu from '@material-ui/core/Menu';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
  Link,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntPriceType } from '../../api/models/EntPriceType'; // import interface PriceType
import { EntNurse } from '../../api/models/EntNurse'; // import interface Nurse
import { EntMedicalfile } from '../../api/models/EntMedicalfile'; // import interface Medicalfile

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },

  menuButton: {
    marginRight: theme.spacing(100),
  },
  title: {
    flexGrow: 1,
  },

}));

interface saveDenExpen {
  PriceType: Number;
  Nurse: Number;
  Medicalfile: Number;
  Added: String;
  
}

const saveDenExpen: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [dentalexpense, setDentalexpenses] = React.useState<Partial<saveDenExpen>>({});
  const [pricetypes, setPriceTypes] = React.useState<EntPriceType[]>([]); //การประกาศตัวแปร โดยที่เราจะดึงมาใช้ แล้ว Ent ได้มาจากการเจน API
  const [nurses, setNurses] = React.useState<EntNurse[]>([]);
  const [medicalfiles, setMedicalfiles] = React.useState<EntMedicalfile[]>([]);

  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  // alert setting
  //const Swal = require('sweetalert2')
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
  });

  // set data to object dentalexpense
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof dentalexpense;
    const { value } = event.target;
    setDentalexpenses({ ...dentalexpense, [name]: value });
    console.log(dentalexpense);
  };
  
  const getPriceType = async () => {
    const res = await http.listPricetype({ limit: 4, offset: 0 });
    setPriceTypes(res);
  };

  
  const getNurse = async () => {
    const res = await http.listNurse({ limit: 4, offset: 0 });
    setNurses(res);
  };

  const getMedicalfile = async () => {
    const res = await http.listMedicalfile({ limit: 4, offset: 0 });
    setMedicalfiles(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getPriceType();
    getMedicalfile();
    getNurse();
    
  }, []);

  // clear input form
  function clear() {
    setDentalexpenses({});
  }

  // function save data
  function save() {
    dentalexpense.Added += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/dentalexpenses';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dentalexpense),
    };

    console.log(dentalexpense); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {console.log(data.save);
        console.log(requestOptions)
        if (data.status == true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  function redirecLogOut() {
    //redirec Page ... http://localhost:3000/
    window.location.href = "http://localhost:3000/";
  }

  function redirecSearch() {
    //redirec Page ... http://localhost:3000/Showmed
    window.location.href = "http://localhost:3000/Showmed";
  }

  return (
    <Page theme={pageTheme.service}>
      <Header
       title="Dental System"
       subtitle="ระบบบันทึกค่ารักษา">

      <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 20 }}>.... ....</div>

      <Button aria-controls="simple-menu" aria-haspopup="true" color="default" onClick={handleClick}>
        Menu
      </Button>
      <Menu
        id="simple-menu"
        anchorEl={anchorEl}
        keepMounted
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >
        <MenuItem onClick={redirecSearch}>SEARCH</MenuItem>
        <MenuItem onClick={redirecLogOut}>LOGOUT</MenuItem>
      </Menu>
      
     </Header>

     
         


      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            
            <Grid item xs={3}>
              <div className={classes.paper}>ประเภทการชำระ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกประเภทการชำระ</InputLabel>
                <Select
                  name="PriceType"
                  value={dentalexpense.PriceType || ''} 
                  onChange={handleChange}
                >
                  {pricetypes.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ข้อมูลการรักษาผู้ป่วย</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกข้อมูลการรักษาผู้ป่วย</InputLabel>
                <Select
                  name="Medicalfile"
                  value={dentalexpense.Medicalfile || ''} 
                  onChange={handleChange}
                >
                  {medicalfiles.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.detail}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่ทำการรักษา</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  
                  label="เลือกวันที่"
                  name="Added"
                  type="datetime-local"
                  value={dentalexpense.Added || ''} 
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>พยาบาล</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกพยาบาลที่ทำการบันทึก</InputLabel>
                <Select
                  name="Nurse"
                  value={dentalexpense.Nurse || ''} 
                  onChange={handleChange}
                >
                  {nurses.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.nurseName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกประวัติทันตกรรม
              </Button>
              &emsp;
              <Link component={RouterLink} to="/Menu">
              <Button
                variant="contained"
                color="default"
                size="large"
                style={{ marginLeft: 5 }}
              >
                กลับ
              </Button>
              </Link>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default  saveDenExpen;