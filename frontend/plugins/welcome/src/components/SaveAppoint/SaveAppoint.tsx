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
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntPatient } from '../../api/models/EntPatient'; // import interface Patient
import { EntDentist } from '../../api/models/EntDentist'; // import interface Dentist
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
    patient: number;
    dentist: number;
    datetime: String;
    detail: String;
    room: number;

}

const Appointment: FC<{}> = () => {
    const classes = useStyles();
    const http = new DefaultApi();
    const [appointment, setAppointment] = React.useState<Partial<appointment>>({});
  
    const [patients, setPatients] = React.useState<EntPatient[]>([]);
    const [dentists, setDentists] = React.useState<EntDentist[]>([]);
    const [rooms, setRooms] = React.useState<EntRoom[]>([]);

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
  
    const getDentists = async () => {
      const res = await http.listDentist({ limit: 10, offset: 0 });
      setDentists(res);
    };
  
    const getRooms = async () => {
      const res = await http.listRoom({ limit: 10, offset: 0 });
      setRooms(res);
    };
  
// Lifecycle Hooks
useEffect(() => {
    getPatients();
    getDentists();
    getRooms();
  }, []);

// set data to object appointment
const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Appointment;
    const { value } = event.target;
    setAppointment({ ...appointment, [name]: value });
    console.log(appointment);
};

// clear input form
function clear() {
    setAppointment({});
}

// function save data
function save() {
    appointment.datetime += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/appointments';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(appointment),
};


    console.log(appointment); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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
      <Header style={HeaderCustom} title={`การนัดหมายผู้ป่วย`}
              subtitle="Welcome to Appointment System.">
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ผู้ป่วย</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกรายชื่อผู้ป่วย</InputLabel>
                <Select
                  name="patient"
                  label = "เลือกรายชื่อผู้ป่วย"
                  value={appointment.patient || ''}
                  onChange={handleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.id}&emsp;
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
                    id="detail"
                    label="กรอกสาเหตุการนัดหมาย"
                    name="detail"
                    variant="outlined"
                    multiline
                    rows={4}
                    type="string"
                    size="medium"
                    value={appointment.detail}
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
                  name="datetime"
                  type="datetime-local"
                  value={appointment.datetime || ''} 
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
                  name="room"
                  label="เลือกห้องตรวจ"
                  value={appointment.room || ''}
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