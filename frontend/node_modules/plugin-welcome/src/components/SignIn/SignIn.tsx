import React, { FC, useState , useEffect } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { DefaultApi, EntNurse } from '../../api';
import { Alert } from '@material-ui/lab'; // alert
import FaceIcon from '@material-ui/icons/Face';

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://github.com/sut63/team03">
          Dental System
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: "#CC3333",
  },
  form: {
    width: '100%',
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const SignIn: FC<{}> = () => {

  const classes = useStyles();
  const api = new DefaultApi();

  const [status, SetStatus] = useState(false);
  const [loading, SetLoading] = useState(true);
  const [alert, SetAlert] = useState(Boolean);

  const [nurse, setNurse] = useState<EntNurse[]>([]);
  const [email, setEmail] = useState(String);
  const [password, setPassword] = useState(String);

  const PasswordhandelChange = (event : any) => {
    setPassword(event.target.value as string);
  }
  const EmailhandelChange = (event : any) => {
    setEmail(event.target.value as string);
  };
  console.log("email",email);


  useEffect(() => {
    const getNurse = async () => {
      const res: any = await api.listNurse({ offset: 0 });
      SetLoading(false);
      setNurse(res);
    };
    getNurse();
    localStorage.clear();
  }, [loading]);


  const SinginhandleChange = async () => {
    nurse.map((item: any) => {
      console.log(item.email);
      if (item.email == email && item.password == password) {

        SetAlert(true);
        localStorage.setItem('nurse-id', JSON.stringify(item.id));
        localStorage.setItem('nurse-name', JSON.stringify(item.name));
        localStorage.setItem('nurse-email', JSON.stringify(item.email));
        history.pushState("", "", "/welcome");
        window.location.reload(false);
      }
    })

  SetStatus(true);
    // const timer = setTimeout(() => {
    //   SetStatus(false);
    // }, 1000);
  };

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      {status ? (
            <div>
              {alert ? (
                <Alert severity="success">Login Succese</Alert>
              ) : (
                <Alert severity="warning" style={{ marginTop: 20 }}>
                  email or password incorrect!!!
                </Alert>
              )}
            </div>
          ) : null}
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <FaceIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign In
        </Typography>
        <form className={classes.form} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
            onChange={EmailhandelChange}
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
            autoComplete="current-password"
            onChange={PasswordhandelChange}
          />
          <Button
            style={{ backgroundColor: '#3333FF' }}
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={() => {
              SinginhandleChange()
            }
            }
          >
            Sign In
          </Button>
        </form>
      </div>
      <Box mt={8}>
        <Copyright />
      </Box>
    </Container>
  );
};

export default SignIn;