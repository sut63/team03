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
  MedicalCare: number;
  Disease: number;

}

const SavePatient: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [patient, setPatient] = React.useState<Partial<savePatient>>({});
  const [patientIDError, setPatientIDError] = React.useState('');
  const [nameError, setnameError] = React.useState('');
  const [cardIDError, setCardIDError] = React.useState('');
  const [telError, setTelError] = React.useState('');
  const [ageError, setAgeError] = React.useState('');
  const [genders, setGenders] = React.useState<EntGender[]>([]);
  const [medicalcares, setMedicalcares] = React.useState<EntMedicalCare[]>([]);
  const [diseases, setDiseases] = React.useState<EntDisease[]>([]);

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 5000,
    timerProgressBar: true,
  });

  const getGender = async () => {
    const res = await http.listGender({ limit: 10, offset: 0 });
    setGenders(res);
  };

  const getMedicalcare = async () => {
    const res = await http.listMedicalcare({ limit: 10, offset: 0 });
    setMedicalcares(res);
  };

  const getDisease = async () => {
    const res = await http.listDisease({ limit: 10, offset: 0 });
    setDiseases(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getMedicalcare();
    getGender();
    getDisease();
  }, []);

  // set data to object patient
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,) => {
    const name = event.target.name as keyof typeof patient;
    const { value } = event.target;
    setPatient({ ...patient, [name]: value });
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    console.log(patient);
  }; 


  //ฟังก์ชั่นสำหรับ validate เลขประจำตัวผู้ป่วย
  const validatepatientID = (val: string) => {
    return val.match("[P]\\d{6}");
  }
  
  //ฟังก์ชั่นสำหรับ validate ชื่อ
  const Validatename = (val: string) => {
    return val.match("");
  } 

  // ฟังก์ชั่นสำหรับ validate เลขประจำตัวประชาชน
  const validatecardID = (val: string) => {
    return val.match("^[0-9]{13}") && val.length <= 13;
    //return val.match("^[0-9]{13}$");
  }

  // ฟังก์ชั่นสำหรับ validate เบอร์โทรศัพท์
  const validatetTel = (val: string) => {
    return val.match("^[0]\\d{9}")&& val.length <= 10 ;
  }

  const ValidateAge = (val: number) => {
    return val <= 200 && val >= 1 ? true : false;
  }

  // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
  const checkPattern = (id: string, value: string) => {
    switch(id) {
      case 'PatientID':
        validatepatientID(value) ? setPatientIDError('') : setPatientIDError('รหัสประจำตัวผู้ป่วยขึ้นต้นด้วย P ตามด้วยตัวเลข 6 ตัว');
        return;
      case 'name' :    
        Validatename(value) ? setnameError('') : setnameError("กรุณากรอกชื่อ-นามสกุล");
        return; 
      case 'CardID':
        validatecardID(value) ? setCardIDError('') : setCardIDError('กรุณากรอกเลขประจำตัวประชาชน 13 หลัก');
        return;
      case 'Tel':
        validatetTel(value) ? setTelError('') : setTelError('กรุณากรอกเบอร์โทรศัพท์ขึ้นต้นด้วย 0 ตามด้วยเลข 9 ตัว');
        return;
      case 'Age':
        ValidateAge(Number(value)) ? setAgeError('') : setAgeError("กรอกอายุตั้งแต่ 1-200 ปีเท่านั้น");
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

  // clear input form
  function clear() {
    setPatient({});
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'PatientID':
        alertMessage("error","รหัสประจำตัวผู้ป่วยขึ้นต้นด้วย P ตามด้วยตัวเลข 6 ตัว");
        return;
      case 'Name':
        alertMessage("error","กรุณากรอกชื่อ-นามสกุล");
        return;
      case 'CardID':
        alertMessage("error","กรุณากรอกเลขประจำตัวประชาชน 13 หลัก");
        return; 
      case 'Tel':
        alertMessage("error","กรุณากรอกเบอร์โทรศัพท์ขึ้นต้นด้วย 0 ตามด้วยเลข 9 ตัว");
        return; 
      case 'Age':
        alertMessage("error","กรอกอายุตั้งแต่ 1-200 ปีเท่านั้น");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  // function save data
  function save() {
    if (patient.Age) {
      var age: number = +patient.Age;
      patient.Age = age;
    }
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
        if (data.status == true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
  }

  return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`บันทึกประวัติผู้ป่วย`}
              subtitle="Dental System - Welcome to Patient">
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>รหัสประจำตัวผู้ป่วย</div>
            </Grid>
            <Grid item xs={6}>
            <FormControl
                  fullWidth
                  className={classes.formControl}
                  variant="outlined">
                  <TextField
                    error={patientIDError ? true : false}
                    label="รหัสประจำตัวผู้ป่วย"
                    name="PatientID"
                    variant="outlined"
                    inputProps={{ maxLength: 7 }}
                    helperText= {patientIDError}
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
                    label="กรอกชื่อ - นามสกุล"
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
                    error={cardIDError ? true : false}
                    label="กรอกเลขบัตรประจำตัวประชาชน"
                    name="CardID"
                    variant="outlined"
                    inputProps={{ maxLength: 13 }}
                    helperText= {cardIDError}
                    size="medium"
                    value={patient.CardID || ''}
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
                  label="เลือกเพศ"
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
                    error={telError ? true : false}
                    label="กรอกเบอร์โทรศัพท์"
                    name="Tel"
                    variant="outlined"
                    inputProps={{ maxLength: 10 }}
                    helperText= {telError}
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
                    error={ageError ? true : false}
                    name="Age"
                    label="อายุ"
                    variant="outlined"
                    size="medium"
                    value={patient.Age || ''}
                    helperText={ageError}
                    onChange={handleChange}
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
                  name="MedicalCare"
                  value={patient.MedicalCare || ''} // (undefined || '') = ''
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

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกข้อมูลประวัติ
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
