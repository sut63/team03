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
  Avatar,
  
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
  QueueID:   String;
	Patient:   Number;
	Phone:     String;
  Dentist:   Number;
  Dental:    String;
  QueueTime: String;

}

const SaveQueue: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [queue, setQueue] = React.useState<Partial<Queue>>({});
  const [dentists, setDentists] = React.useState<EntDentist[]>([]);
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  
  const [QueueIDError, setQueueIDError] = React.useState('');
  const [PhoneError, setPhoneError] = React.useState('');
  const [DentalError, setDentalError] = React.useState('');
  

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
    event: React.ChangeEvent<{ name?: string; value: any }>,  ) => {
    const name = event.target.name as keyof typeof queue;
    const { value } = event.target;
    setQueue({ ...queue, [name]: value });
    const validateValue = value.toString()
    checkPattern(name, validateValue)
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

  
  // ฟังก์ชั่นสำหรับ validate รหัสคิว
  const validateQueueID = (val: string) => {
    return val.match("[Q]\\d{3}") && val.length <= 4;
  }

  // ฟังก์ชั่นสำหรับ validate โทรศัพท์
  const validatePhone = (val: string) => {
    return val.length == 10 ? true : false;
  }
 
  // ฟังก์ชั่นสำหรับ validate ทันตกรรม
  const validateDental = (val: string) => {
    return val.length > 30 ? false : true;
  }

  // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
  const checkPattern  = (id: string, value: string) => {
    switch(id) {
      case 'QueueID':
        validateQueueID(value) ? setQueueIDError('') : setQueueIDError('รหัสคิวขึ้นต้นด้วย Q ตามด้วยตัวเลข 3 ตัว');
        return;
      case 'Phone':
        validatePhone(value) ? setPhoneError('') : setPhoneError('หมายเลขโทรศัพท์ต้องมี 10 หลัก');
        return;
        case 'Dental':
        validateDental(value) ? setDentalError('') : setDentalError('มีตัวอักษรเกิน 30 ตัวอักษร');
        return;
      default:
        return;
    }
  }

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  //กำหนดข้อความ error
  const checkCaseSaveError = (s: string) => {
    switch(s) {
      case 'QueueID':
        alertMessage("error", "รหัสคิวขึ้นต้นด้วย Q ตามด้วยตัวเลข 3 ตัว");
        return;
      case 'Phone':
        alertMessage("error", "หมายเลขโทรศัพท์ต้องมี 10 หลัก");
        return;
      case 'Dental':
        alertMessage("error", "มีตัวอักษรเกิน 30 ตัวอักษร");
        return;
      default:
        alertMessage("error", " บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  

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
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
  };

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
              <div className={classes.paper}>รหัสคิว</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                error = {QueueIDError ? true : false}
                label="รหัสคิว"
                variant="outlined"
                name="QueueID"
                type="string" 
                helperText={QueueIDError}
                value={queue.QueueID || ''}
                onChange={handleChange}
                />
              </FormControl>
            </Grid>

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
              <div className={classes.paper}>หมายเลขโทรศัพท์ติดต่อฉุกเฉิน</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                error = {PhoneError ? true : false}
                name="Phone"
                label="หมายเลขโทรศัพท์ติดต่อฉุกเฉิน" 
                helperText={PhoneError}
                variant="outlined"
                type="string"
                value={queue.Phone || ''}
                onChange={handleChange}
                />

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
                error = {DentalError ? true : false}
                name="Dental"
                label="ทันตกรรม" 
                helperText={DentalError}
                variant="outlined"
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