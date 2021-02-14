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
import { EntMedicalfile } from '../../api/models/EntMedicalfile'; // import interface Dentist
import { EntNurse } from '../../api/models/EntNurse'; // import interface Nurse
import { EntPricetype } from '../../api/models/EntPricetype'; // import interface Patient


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

interface saveDenExpen {
  Medicalfile: Number;
  Pricetype: Number;
  Nurse: Number;
  Amount: String;
  AddedTime: String;
  Tax: String;
  Name: String;
  Phone: String;
}

const SaveDenExpen: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [dentalexpense, setDentalexpense] = React.useState<Partial<saveDenExpen>>({});
  const [medicalfiles, setMedicalfiles] = React.useState<EntMedicalfile[]>([]); //การประกาศตัวแปร โดยที่เราจะดึงมาใช้ แล้ว Ent ได้มาจากการเจน API
  const [pricetypes, setPricetypes] = React.useState<EntPricetype[]>([]);
  const [nurses, setNurses] = React.useState<EntNurse[]>([]);
  const [TaxError, setTaxError] = React.useState('');
  const [NameError, setNameError] = React.useState('');
  const [PhoneError, setPhoneError] = React.useState('');
  
 

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
    getMedicalfile();
    getPricetype();
    getNurse();
  }, []);

  // set data to object medicalfile
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,  ) => {
    const name = event.target.name as keyof typeof dentalexpense;
    const { value } = event.target;
    setDentalexpense({ ...dentalexpense, [name]: value });
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    console.log(dentalexpense);
  };
  
  
  const getPricetype = async () => {
    const res = await http.listPricetype({ limit: 10, offset: 0 });
    setPricetypes(res);
  };

  const getMedicalfile = async () => {
    const res = await http.listMedicalfile({ limit: 10, offset: 0 });
    setMedicalfiles(res);
  };

  const getNurse = async () => {
    const res = await http.listNurse({ limit: 10, offset: 0 });
    setNurses(res);
  };

   
   const validateTax = (val: string) => {
    return val.match("[R]\\d{10}") && val.length <= 11;
  }

  
  const validateName = (val: string) => {
    return val.match("^[ก-๏\\s]+$"); 
  }

  
  const validatePhone = (val: string) => {
    return val.match("^[0-9]{10}$");
  }
  
  
  const checkPattern  = (id: string, value: string) => {
    switch(id) {
      case 'Tax':
        validateTax(value) ? setTaxError('') : setTaxError('รูปแบบหมายเลขกำกับภาษีไม่ถูกต้อง');
        return;
      case 'Name':
        validateName(value) ? setNameError('') : setNameError('กรอกชื่อเป็นภาษาไทยเท่านั้น');
        return;
      case 'Phone':
        validatePhone(value) ? setPhoneError('') : setPhoneError('กรอกเบอร์โทรศัพท์10ตัวเท่านั้น');
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
  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'Tax':
        alertMessage("error", "รูปแบบหมายเลขกำกับภาษีไม่ถูกต้อง");
        return;
      case 'Name':
        alertMessage("error", "กรุณากรอกเฉพาะภาษาไทย");
        return;
      case 'Phone':
        alertMessage("error", " กรอกเบอร์โทรศัพท์10ตัวเท่านั้น");
        return;
      case 'Amount':
        alertMessage("error","กรุณากรอกเป็นตัวเลขและมีค่าไม่น้อยกว่า 0"); 
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  

  // clear input form
  function clear() {
    setDentalexpense({});
  }

  // function save data
  function save() {
    dentalexpense.AddedTime += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/dentalexpenses';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dentalexpense),
    };

    console.log(dentalexpense); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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
       subtitle="ระบบบันทึกค่ารักษา">
     </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เลขที่กำกับภาษี</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                error = {TaxError ? true : false}
                label="ระบุเลขที่กำกับภาษี" 
                variant="outlined" 
                name="Tax"
                type="string"
                helperText= {TaxError}
                value={dentalexpense.Tax || ''}
                onChange={handleChange}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>บริการทันตกรรม</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกข้อมูลการรับบริการทันตกรรม</InputLabel>
                <Select
                  name="Medicalfile"
                  value={dentalexpense.Medicalfile || ''} 
                  onChange={handleChange}
                >
                  {medicalfiles.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.id}&emsp;
                        {item.detial}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ค่าบริการ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
              
              <TextField
              id="Amount"
              name = "amount"
              label="ค่าบริการ" 
              variant="outlined"
              
              value={dentalexpense.Amount}
              
              
              onChange={handleChange}
              style={{ width: 300 }}
              />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ประเภทการชำระ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกประเภทการชำระ</InputLabel>
                <Select
                  name="Pricetype"
                  value={dentalexpense.Pricetype || ''} 
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
              <div className={classes.paper}>ชื่อผู้ชำระค่าบริการ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                error = {NameError ? true : false}
                helperText={NameError}
                label="ชื่อผู้ชำระค่าบริการ" 
                variant="outlined" 
                name="Name"
                type="string"
                value={dentalexpense.Name || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เบอร์ที่ติดต่อได้</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                error = {PhoneError ? true : false}
                helperText={PhoneError}
                label="ระบุเบอร์โทรศัพท์" 
                variant="outlined" 
                name="Phone"
                type="string"
                value={dentalexpense.Phone || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

           

            

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่ทำการชำระ</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  
                  label="เลือกวันที่"
                  name="AddedTime"
                  type="datetime-local"
                  value={dentalexpense.AddedTime || ''} 
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
                บันทึกค่ารักษา
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

export default  SaveDenExpen;