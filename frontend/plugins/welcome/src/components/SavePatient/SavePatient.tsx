import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Link as RouterLink } from 'react-router-dom';
import Swal from 'sweetalert2'; // alert


import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Link,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntGender } from '../../api/models/EntGender'; // import interface Gender
import { EntMedicalCare } from '../../api/models/EntMedicalCare'; // import interface MedicalCare
import { EntNurse } from '../../api/models/EntNurse'; // import interface Nurse
import { EntDisease } from '../../api/models/EntDisease'; // import interface Disease

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

interface savePatient {
  PatientID: String;
  Name: String;
  CardID: String;
  Gender: number;
  Tel: String;
  Birthday: String;
  Age: number;
  Medicalcare: number;
  Nurse: number;
  Disease: number;

}

const SavePatient: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [patient, setPatient] = React.useState<Partial<savePatient>>({});
  const [genders, setGenders] = React.useState<EntGender[]>([]);
  const [medicalcares, setMedicalcares] = React.useState<EntMedicalCare[]>([]);
  const [nurses, setNurses] = React.useState<EntNurse[]>([]);
  const [diseases, setDiseases] = React.useState<EntDisease[]>([]);

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
  });

  const getGender = async () => {
    const res = await http.listGender({ limit: 2, offset: 0 });
    setGenders(res);
  };

  const getMedicalcare = async () => {
    const res = await http.listMedicalcare({ limit: 4, offset: 0 });
    setMedicalcares(res);
  };

  const getNurse = async () => {
    const res = await http.listNurse({ limit: 2, offset: 0 });
    setNurses(res);
  };

  const getDisease = async () => {
    const res = await http.listDisease({ limit: 10, offset: 0 });
    setDiseases(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getMedicalcare();
    getGender();
    getNurse();
    getDisease();
  }, []);

  // set data to object patient
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof patient;
    const { value } = event.target;
    setPatient({ ...patient, [name]: value });
    console.log(patient);
  };

  const handleNumberChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof patient;
    const { value } = event.target;
    setPatient({ ...patient, [name]: + value });
    console.log(patient);
  };

  // clear input form
  function clear() {
    setPatient({});
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/patients';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(patient),
    };

    console.log(JSON.stringify(patient)); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {console.log(data.save)
        console.log(requestOptions)
        if (data.status === true) {
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
      <Header style={HeaderCustom} title={`บันทึกประวัติผู้ป่วย`}>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เลขประจำตัวผู้ป่วย</div>
            </Grid>
            <Grid item xs={6}>
            <FormControl
                  fullWidth
                  className={classes.formControl}
                  variant="outlined">
                  <TextField
                    label="เลขประจำตัวผู้ป่วย"
                    name="PatientID"
                    variant="outlined"
                    size="medium"
                    value={patient.PatientID || ''}
                    onChange={handleChange}
                  />
                </FormControl>
            </Grid>
            
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อ - นามสกุล</div>
            </Grid>
            <Grid item xs={6}>
            <FormControl
                  fullWidth
                  className={classes.formControl}
                  variant="outlined">
                  <TextField
                    label="ชื่อ - นามสกุล"
                    name="Name"
                    variant="outlined"
                    size="medium"
                    value={patient.Name || ''}
                    onChange={handleChange}
                  />
                </FormControl>
            </Grid>
            
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เลขบัตรประจำตัวประชาชน</div>
            </Grid>
            <Grid item xs={6}>
            <FormControl
                  fullWidth
                  className={classes.formControl}
                  variant="outlined">
                  <TextField
                    label="เลขบัตรประจำตัวประชาชน"
                    name="Name"
                    variant="outlined"
                    size="medium"
                    value={patient.Name || ''}
                    onChange={handleChange}
                  />
                </FormControl>
            </Grid>

            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เพศ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เพศ</InputLabel>
                <Select
                  name="Gender"
                  label="เพศ"
                  value={patient.Gender || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {genders.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เบอร์โทรศัพท์</div>
            </Grid>
            <Grid item xs={6}>
            <FormControl
                  fullWidth
                  className={classes.formControl}
                  variant="outlined">
                  <TextField
                    label="เบอร์โทรศัพท์"
                    name="Tel"
                    variant="outlined"
                    size="medium"
                    value={patient.Tel || ''}
                    onChange={handleChange}
                  />
                </FormControl>
            </Grid>
         
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>วันเดือนปีเกิด</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="วันเดือนปีเกิด"
                  name="Birthday"
                  type="date"
                  value={patient.Birthday || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>อายุ</div>
            </Grid>
            <Grid item xs={6}>
            <FormControl
                  fullWidth
                  className={classes.formControl}
                  variant="outlined">
                  <TextField
                    label="อายุ"
                    name="Age"
                    type="Number"
                    variant="outlined"
                    size="medium"
                    value={patient.Age || ''}
                    onChange={handleNumberChange}
                  />
                </FormControl>
            </Grid>

            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>โรคประจำตัว</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>โรคประจำตัว</InputLabel>
                <Select
                  label="โรคประจำตัว"
                  name="Disease"
                  value={patient.Disease || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {diseases.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>สิทธิในการรักษา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>สิทธิในการรักษา</InputLabel>
                <Select
                  label="สิทธิในการรักษา"
                  name="Medicalcare"
                  value={patient.Medicalcare || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {medicalcares.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>พยาบาล</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>พยาบาลผู้บันทึก</InputLabel>
                <Select
                  label="พยาบาลผู้บันทึก"
                  name="Nurse"
                  value={patient.Nurse || ''} // (undefined || '') = ''
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
                บันทึกข้อมูล
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

export default SavePatient;
