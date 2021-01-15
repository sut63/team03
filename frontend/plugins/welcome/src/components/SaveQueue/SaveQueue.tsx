import React, { FC, useEffect } from 'react';
import { makeStyles, } from '@material-ui/core/styles';
import { ContentHeader,Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  
} from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';

import Swal from 'sweetalert2'; // alert
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntPatient } from '../../api/models/EntPatient'; 
import { EntDentist } from '../../api/models/EntDentist'; 


// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  margin: {
    margin: theme.spacing(2),
 },
 insideLabel: {
  margin: theme.spacing(1),
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
}));

interface Queue {
  Patient: Number;
  Dentist: Number;
  Dental: String;
  QueueTime: String;

}

const SaveQueue: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [queue, setQueue] = React.useState<Partial<Queue>>({});
  const [dentists, setDentists] = React.useState<EntDentist[]>([]);
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  

  // alert setting
  //const Swal = require('sweetalert2')
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
  });

  // set data to object queue
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof queue;
    const { value } = event.target;
    setQueue({ ...queue, [name]: value });
    console.log(queue);
  };
  

  const getPatient = async () => {
    const res = await http.listPatient({ limit: 5, offset: 0 });
    setPatients(res);
  };

  const getDentist = async () => {
    const res = await http.listDentist({ limit: 5, offset: 0 });
    setDentists(res);
  };


  // Lifecycle Hooks
  useEffect(() => {
    getPatient();
    getDentist();
  }, []);

  // clear input form
  function clear() {
    setQueue({});
  }

  // function save data
  function save() {
    queue.QueueTime += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/queues';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(queue),
    };

    console.log(queue); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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
       title="ระบบจองคิวผู้ป่วย"
       subtitle="Queue Order System">
         <Button variant="contained" color="default" href="/SignInNurse" startIcon={<ExitToAppRoundedIcon />}> Logout
           </Button>
     </Header>
      <Content>
        <ContentHeader title="จองคิวผู้ป่วย"> 
              <Button onClick={save} variant="contained"  color="primary" startIcon={<SaveIcon />}>บันทึกการจองคิว </Button>
              <Button style={{ marginLeft: 20 }} component={RouterLink} to="/Menu" variant="contained" >กลับ</Button>
        </ContentHeader>

        <Container maxWidth="sm">
        
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>ผู้ป่วยที่ต้องการจองคิว</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกผู้ป่วย</InputLabel>
                <Select
                  name="Patient"
                  value={queue.Patient || ''} 
                  onChange={handleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>


            <Grid item xs={4}>
              <div className={classes.paper}>ทันตแพทย์ที่จะทำการรักษา</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกทันตแพทย์</InputLabel>
                <Select
                  name="Dentist"
                  value={queue.Dentist || ''} 
                  onChange={handleChange}
                >
                  {dentists.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            
            <Grid item xs={4}>
              <div className={classes.paper}>ทันตกรรมที่ต้องการทำ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                
                label="ทันตกรรม" 
                variant="outlined" 
                name="Dental"
                type="string"
                value={queue.Dental || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>เวลา</div>
            </Grid>
            <Grid item xs={8}>
              <form className={classes.container} noValidate>
                <TextField
                  
                  label="เวลา"
                  name="QueueTime"
                  type="datetime-local"
                  value={queue.QueueTime || ''} 
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>
            
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default  SaveQueue;