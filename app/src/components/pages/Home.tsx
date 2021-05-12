import React, { useEffect, useState } from 'react';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import Timeline from '@material-ui/lab/Timeline';
import TextField from '@material-ui/core/TextField';
import TimelineParts from '../organisms/TimelineParts';
import Header from '../organisms/Header';
import axios from 'axios'

const useStyles = makeStyles((theme) => ({
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(1, 0, 5),
  },
}));

type Diarys = {
  body: string,
  create_at: string
}[]

const Home: React.FC = () => {

  const classes = useStyles();
  const [diary, setDiary] = useState<Diarys>([])
  const [talk, setTalk] = useState('')

  useEffect(() => {

    axios.get<Diarys>('http://localhost:8082/diary', {
      xsrfHeaderName: 'X-CSRF-Token',
      withCredentials: true
    }).then((respons) => {
      const d = respons.data.map((c) => ({ body: c.body, create_at: c.create_at }))
      console.log(respons);

      setDiary(d)
    })

  }, [])
  return (

    <>

      <Header />

      <form className={classes.form} noValidate
        onSubmit={async (e) => {
          e.preventDefault()

          // ログイン処理
          await axios.post('http://localhost:8082/diary', {
            body: talk,
          }, {
            xsrfHeaderName: 'X-CSRF-Token',
            withCredentials: true
          }).then((e) => console.log(e));

          setDiary([{ body: talk, create_at: Date() }, ...diary])
          setTalk('')

        }}>
        <TextField
          variant="outlined"
          margin="normal"
          fullWidth
          id="word"
          label="今の状況を一言で。"
          name="word"
          autoComplete="word"
          autoFocus
          value={talk}
          onChange={(e) => { setTalk(e.target.value) }}
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

      <Timeline align="alternate">

        {
          diary.map((d, index) => (
            <TimelineParts key={index} body={d.body} />
          ))
        }

      </Timeline>
    </>
  );
}

export default Home