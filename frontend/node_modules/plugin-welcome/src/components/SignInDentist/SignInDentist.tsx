import React, { FC, useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import {Header,Page,pageTheme,} from '@backstage/core';
import {TextField,Button,Grid,Link,} from '@material-ui/core';
import { Alert } from '@material-ui/lab'; // alert
import { DefaultApi } from '../../api/apis';
import { EntDentist } from '../../api/models/EntDentist';

const HeaderCustom = {
  minHeight: '120px',
};

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  form: {
    width: '100%',
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));


const SignInDentist: FC<{}> = () => {
  const classes = useStyles();
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(Boolean);
  const api = new DefaultApi();
  const [dentists, setDentist] = useState<EntDentist[]>([]);
  const [emails, setEmail] = React.useState(String);
  const [password, setPassword] = React.useState(String);


  const emailHandleChange = (event: any) => {
    setEmail(event.target.value as string);
  };
  const passwordHandleChange = (event: any) => {
    setPassword(event.target.value as string);
  };

  const getDentist = async () => {
    const res: any = await api.listDentist({ limit: 0, offset: 0 });
    setDentist(res);
  };

  const resetData = async () => {
    localStorage.clear();
  };
  useEffect(() => {
    getDentist();
    resetData();
  }, []);

  const SigninCheck = async () => {
    var checkDentist = false;

    dentists.map((item: any) => {
      console.log(item.email);
      if (item.email == emails && item.password == password) {
        setAlert(true);
        checkDentist = true;
        localStorage.setItem('dentist-id', JSON.stringify(item.id));
        localStorage.setItem('dentist-name', JSON.stringify(item.name));
        history.pushState('', '', '/SaveAppoint'); /////////มาแก้ที่หลังไปยังหน้ารวมเภสัชกร
        window.location.reload(false);
      }
    });

    setStatus(true);
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    <div className={classes.paper}>
      <Page theme={pageTheme.website}>
        <Header style={HeaderCustom} title={`Doctor Information`}
          subtitle="กรุณาบันทึกข้อมูลก่อนเข้าสู่ระบบ">
            {status ? (
            <div>
              {alert ? (
                <Alert severity="success">เข้าสู่ระบบสำเร็จ</Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    ไม่พบข้อมูลในระบบ
                  </Alert>
                )}
            </div>
          ) : null}
        </Header>
        <form className={classes.submit} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="email"
            name="email"
            autoComplete="email"
            autoFocus
            onChange={emailHandleChange}
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            onChange={passwordHandleChange}
            autoComplete="current-password"
            
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={SigninCheck}
            >
            Sign In
          </Button>
          <Grid container>
            <Grid item>
              <Link href="/user" variant="body2">
                {"Don't have an account? Sign Up"}
              </Link>
            </Grid>
          </Grid>
        </form>
      </Page>
    </div>
  );
};
export default SignInDentist;