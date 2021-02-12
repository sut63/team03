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
  Button,
  Link,
} from '@material-ui/core';

import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntPatient } from '../../api/models/EntPatient'; // import interface Patient
import { EntRoom } from '../../api/models/EntRoom'; // import interface Room

// header css
const HeaderCustom = {
    minHeight: '50px',
  };
  
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
  }));

interface appointment {
    AppointID: String;
    Patient: Number;
    Dentist: Number;
    Datetime: String;
    Detail: String;
    Room: Number;
    Remark: String;

}

const Appointment: FC<{}> = () => {
    const classes = useStyles();
    const http = new DefaultApi();
    const [appointment, setAppointment] = React.useState<Partial<appointment>>({});
  
    const [patients, setPatients] = React.useState<EntPatient[]>([]);
    const [rooms, setRooms] = React.useState<EntRoom[]>([]);
    const [DetailError, setDetailError] = React.useState('');
    const [AppointIDError, setAppointIDError] = React.useState('');
    const [RemarkError, setRemarkError] = React.useState('');

// alert setting
const Toast = Swal.mixin({
    position: 'center',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    
});

    const getPatients = async () => {
      const res = await http.listPatient({ limit: 10, offset: 0 });
      setPatients(res);
    };

    const getRooms = async () => {
      const res = await http.listRoom({ limit: 10, offset: 0 });
      setRooms(res);
    };
  
// Lifecycle Hooks
useEffect(() => {
    getPatients();
    getRooms();
  }, []);

// set data to object appointment
const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>, ) => {
    const name = event.target.name as keyof typeof appointment;
    const { value } = event.target;
    setAppointment({ ...appointment, [name]: value });
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    console.log(appointment);
};

 // ฟังก์ชั่นสำหรับ validate รหัสการนัดหมาย
 const validateAppointID = (val: string) => {
  return val.match("[A]\\d{5}")
}

// ฟังก์ชั่นสำหรับ validate สาเหตุการนัดหมาย
const validateDetail = (val: string) => {
  return val.length < 5 ?  false : true;
}

// ฟังก์ชั่นสำหรับ validate หมายเหตุ
const validateRemark = (val: string) => {
  return val.length > 30 ?  false : true;
}

// สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
const checkPattern  = (id: string, value: string) => {
  switch(id) {
    case 'AppointID':
      validateAppointID(value) ? setAppointIDError('') : setAppointIDError('รหัสการนัดหมายขึ้นต้นด้วย A ตามด้วยตัวเลข 5 ตัว');
      return;

    case 'Detail':
      validateDetail(value) ? setDetailError('') : setDetailError('สาเหตุการนัดหมายจำนวนตัวอักษรต่ำกว่า 5 ตัวอักษร');
      return;

    case 'Remark':
      validateRemark(value) ? setRemarkError('') : setRemarkError('หมายเหตุจำนวนตัวอักษรมากกว่า 30 ตัวอักษร');
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
    case 'AppointID':
      alertMessage("error", "รหัสการนัดหมายขึ้นต้นด้วย A ตามด้วยตัวเลข 5 ตัว");
      return;
    case 'Detail':
      alertMessage("error", " สาเหตุการนัดหมายจำนวนตัวอักษรต่ำกว่า 5 ตัวอักษร");
      return;
    case 'Remark':
      alertMessage("error", " หมายเหตุจำนวนตัวอักษรมากกว่า 30 ตัวอักษร");
      return;
    default:
      alertMessage("error", " บันทึกข้อมูลไม่สำเร็จ");
      return;
  }
}


// clear input form
function clear() {
    setAppointment({});
}

// function save data
function save() {
    appointment.Datetime += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/appointments';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(appointment),
};


    console.log(appointment); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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
      <Header style={HeaderCustom} title={`การนัดหมายผู้ป่วย`}
              subtitle="Welcome to Appointment System.">
      <Button variant="contained" color="default" href="/" startIcon={<ExitToAppRoundedIcon />}> Logout
      </Button>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>รหัสการนัดหมาย</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                error = {AppointIDError ? true : false}
                helperText={AppointIDError}
                label="รหัสการนัดหมาย" 
                variant="outlined" 
                name="AppointID"
                type="string"
                value={appointment.AppointID || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>รหัสผู้ป่วย</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกรหัสผู้ป่วย</InputLabel>
                <Select
                  name="Patient"
                  label = "เลือกรายชื่อผู้ป่วย"
                  value={appointment.Patient || ''}
                  onChange={handleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.patientID}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>ผู้ป่วย</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>รายชื่อผู้ป่วย</InputLabel>
                <Select
                  name="Patient"
                  label = "รายชื่อผู้ป่วย"
                  value={appointment.Patient || ''}
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

            <Grid item xs={3}>
              <div className={classes.paper}>สาเหตุการนัดหมาย</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                error = {DetailError ? true : false}
                helperText={DetailError}
                label="สาเหตุการนัดหมาย" 
                variant="outlined" 
                name="Detail"
                type="string"
                value={appointment.Detail || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่นัดหมาย</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  
                  label="เลือกวันที่"
                  name="Datetime"
                  type="datetime-local"
                  value={appointment.Datetime || ''} 
                  className={classes.formControl}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>ห้องตรวจ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกห้องตรวจ</InputLabel>
                <Select
                  name="Room"
                  label="เลือกห้องตรวจ"
                  value={appointment.Room || ''}
                  onChange={handleChange}
                >
                  {rooms.map(item => {
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
              <div className={classes.paper}>หมายเหตุ</div>
            </Grid>
            <Grid item xs={9}>
              <TextField
                error = {RemarkError ? true : false}
                className={classes.formControl}
                id="Remark"
                name="Remark"
                label="หมายเหตุ"
                helperText= {RemarkError}
                value={appointment.Remark || ''}
                variant="outlined"
                onChange={handleChange}
              />
            </Grid>


            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                style={{ marginRight: 5 }}
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกการนัดหมาย
              </Button>

              <Link component={RouterLink} to="/signindentist">
              <Button
                variant="contained"
                color="primary"
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

export default Appointment;