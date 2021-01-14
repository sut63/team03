import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import { Link as RouterLink } from 'react-router-dom';

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
import { EntNurse } from '../../api/models/EntNurse'; // import interface Nurse


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

interface SaveNurse {
  NurseName: String;
  NurseAge: number;
  Email: String;
  Password: String;
}

const SaveNurse: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [nurse, setNurse] = React.useState<Partial<SaveNurse>>({});

  // alert setting
  //const Swal = require('sweetalert2')
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
  });

  // set data to object nurse
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof nurse;
    const { value } = event.target;
    setNurse({ ...nurse, [name]: value });
    console.log(nurse);
  };

  const handleNumberChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof nurse;
    const { value } = event.target;
    setNurse({ ...nurse, [name]: + value });
    console.log(nurse);
  };


  // clear input form
  function clear() {
    setNurse({});
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/nurses';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(nurse),
    };

    console.log(nurse); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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

  return (
    <Page theme={pageTheme.service}>
      <Header
       title="Dental System"
       subtitle="บันทึกข้อมูลพยาบาล">
     </Header>
      <Content>
        <Container maxWidth="sm">
         
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อ-สกุล</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                
                label="ชื่อ-สกุล พยาบาล" 
                variant="outlined" 
                name="nursename"
                type="string"
                value={nurse.NurseName }
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>อายุ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                
                label="อายุ" 
                variant="outlined" 
                name="nurseage"
                type="string"
                value={nurse.NurseAge }
                onChange={handleNumberChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>Email</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                
                label="Email" 
                variant="outlined" 
                name="email"
                type="string"
                value={nurse.Email }
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>Password</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                
                label="Password" 
                variant="outlined" 
                name="password"
                type="string"
                value={nurse.Password }
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

<br></br>
            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกพยาบาล
              </Button>
              &emsp;
              <Link component={RouterLink} to="/signinnurse">
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
         
        </Container>
      </Content>
    </Page>
  );
};

export default  SaveNurse;