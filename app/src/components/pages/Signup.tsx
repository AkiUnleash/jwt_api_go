import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import axios from 'axios'
import { browserHistory } from '../../history'

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.info.dark,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const SignUp: React.FC = () => {

  const classes = useStyles();

  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <AccountCircleIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          ユーザー登録
        </Typography>

        <form className={classes.form} noValidate
          onSubmit={(async (e) => {
            e.preventDefault()

            // サインアップ処理
            await axios.post('http://localhost:8082/account/signup', {
              email: email,
              password: password
            }).then((e) => console.log(e));

            // ログイン処理
            await axios.post('http://localhost:8082/account/login', {
              email: email,
              password: password
            }, {
              xsrfHeaderName: 'X-CSRF-Token',
              withCredentials: true
            }).then((e) => console.log(e));

            browserHistory.push('home')

          })} >
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
            value={email}
            onChange={(e) => setEmail(e.target.value)}
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
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
          >
            登録
          </Button>
        </form>
        <Grid container justify="center">
          <Grid item>
            <Link to="/login">
              {"ログインはこちら"}
            </Link>
          </Grid>
        </Grid>
      </div>
    </Container >
  );
}

export default SignUp