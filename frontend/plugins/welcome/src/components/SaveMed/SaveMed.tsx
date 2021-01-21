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
import { EntDentist } from '../../api/models/EntDentist'; // import interface Dentist
import { EntNurse } from '../../api/models/EntNurse'; // import interface Nurse
import { EntPatient } from '../../api/models/EntPatient'; // import interface Patient


// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
    textAlign: 'right',
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

interface saveMed {
  Dentist: Number;
  Patient: Number;
  Nurse: Number;
  AddedTime: String;
  Detial: String;
  Medno: String;
  DrugAllergy: String;
}

const SaveMed: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [medicalfile, setMedicalfile] = React.useState<Partial<saveMed>>({});
  const [dentists, setDentists] = React.useState<EntDentist[]>([]); //การประกาศตัวแปร โดยที่เราจะดึงมาใช้ แล้ว Ent ได้มาจากการเจน API
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  const [nurses, setNurses] = React.useState<EntNurse[]>([]);
  const [MednoError, setMednoError] = React.useState('');
  const [DetialError, setDetialError] = React.useState('');
  const [DrugAllergyError, setDrugAllergyError] = React.useState('');

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    
  });

  // Lifecycle Hooks
  useEffect(() => {
    getPatient();
    getDentist();
    getNurse();
  }, []);

  // set data to object medicalfile
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,  ) => {
    const name = event.target.name as keyof typeof medicalfile;
    const { value } = event.target;
    setMedicalfile({ ...medicalfile, [name]: value });
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    console.log(medicalfile);
  };
  
  const getPatient = async () => {
    const res = await http.listPatient({ limit: 10, offset: 0 });
    setPatients(res);
  };

  const getDentist = async () => {
    const res = await http.listDentist({ limit: 10, offset: 0 });
    setDentists(res);
  };

  const getNurse = async () => {
    const res = await http.listNurse({ limit: 10, offset: 0 });
    setNurses(res);
  };

   // ฟังก์ชั่นสำหรับ validate รหัสประวัติทันตกรรม
   const validateMedno = (val: string) => {
    return val.match("[M]\\d{3}") && val.length <= 4;
  }

  // ฟังก์ชั่นสำหรับ validate การรักษา
  const validateDetial = (val: string) => {
    return val.length > 30 ? false : true;
  }

  // ฟังก์ชั่นสำหรับ validate การแพ้ยา
  const validateDrugAllergy = (val: string) => {
    return val.length > 30 ? false : true;
  }

  // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
  const checkPattern  = (id: string, value: string) => {
    switch(id) {
      case 'Medno':
        validateMedno(value) ? setMednoError('') : setMednoError('รหัสประวัติทันตกรรมขึ้นต้นด้วย M ตามด้วยตัวเลข 3 ตัว');
        return;
      case 'Detial':
        validateDetial(value) ? setDetialError('') : setDetialError('ห้ามเกิน 30 ตัวอักษร');
        return;
        case 'DrugAllergy':
        validateDrugAllergy(value) ? setDrugAllergyError('') : setDrugAllergyError('ห้ามเกิน 30 ตัวอักษร');
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
      case 'Medno':
        alertMessage("error", "รูปแบบรหัสประวัติทันตกรรมไม่ถูกต้อง");
        return;
      case 'Detial':
        alertMessage("error", " จำนวนตัวอักษรเกิน 30 ตัวอักษร");
        return;
      case 'DrugAllergy':
        alertMessage("error", " จำนวนตัวอักษรเกิน 30 ตัวอักษร");
        return;
      default:
        alertMessage("error", " บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  

  // clear input form
  function clear() {
    setMedicalfile({});
  }

  // function save data
  function save() {
    medicalfile.AddedTime += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/medicalfiles';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(medicalfile),
    };

    console.log(medicalfile); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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
       title="Dental System"
       subtitle="ระบบบันทึกประวัติทันตกรรม">
     </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>รหัสประวัติทันตกรรม</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                error = {MednoError ? true : false}
                label="รหัสประวัติทันตกรรม" 
                variant="outlined" 
                name="Medno"
                type="string"
                helperText= {MednoError}
                value={medicalfile.Medno || ''}
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
                  value={medicalfile.Patient || ''} 
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
              <div className={classes.paper}>ทันตกรรม</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                error = {DetialError ? true : false}
                helperText={DetialError}
                label="ประเภททันตกรรม" 
                variant="outlined" 
                name="Detial"
                type="string"
                value={medicalfile.Detial || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>การแพ้ยา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                error = {DrugAllergyError ? true : false}
                helperText={DrugAllergyError}
                label="การแพ้ยา" 
                variant="outlined" 
                name="DrugAllergy"
                type="string"
                value={medicalfile.DrugAllergy || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ทันตแพทย์</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกทันตแพทย์ที่ทำการรักษา</InputLabel>
                <Select
                  name="Dentist"
                  value={medicalfile.Dentist || ''} 
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

            <Grid item xs={3}>
              <div className={classes.paper}>สิทธิในการรักษา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกสิทธิในการรักษา</InputLabel>
                <Select
                  name="Patient"
                  value={medicalfile.Patient || ''} 
                  onChange={handleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem value={item.id}>{item.edges?.medicalcare?.name}</MenuItem>
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
                  name="AddedTime"
                  type="datetime-local"
                  value={medicalfile.AddedTime || ''} 
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
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

export default  SaveMed;