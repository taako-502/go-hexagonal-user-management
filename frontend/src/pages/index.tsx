import { User } from '@/type/user.type'
import axios from 'axios'
import { useState } from 'react'

export default function Home() {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [error, setError] = useState('')

  const save = async () => {
    const body: User = {
      Username: username,
      Email: email,
    }
    try {
      axios.defaults.baseURL = process.env.NEXT_PUBLIC_BASE_URL
      await axios.post('/user', body)
      setUsername('')
      setEmail('')
      setError('')
    } catch (error) {
      console.error(error)
      if (error instanceof Error) {
        console.error(error.message)
        setError(error.message)
      }
    }
  }

  const handleSubmit = (event: any) => {
    event.preventDefault() // フォーム送信をキャンセル
    console.log('ボタンがクリックされましたが、送信はキャンセルされました')
    // ここにその他の処理を追加
  }

  return (
    <main className="flex min-h-screen flex-col items-center p-24">
      <h1 className="text-2xl">ユーザ管理画面</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="username" className="mr-[2em]">
            ユーザ名:
          </label>
          <input
            id="username"
            type="text"
            className="border bg-gray-600 border-gray-400 rounded-md p-2 mt-[8px]"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div>
          <label htmlFor="email" className="mr-[2em]">
            メールアドレス:
          </label>
          <input
            id="email"
            type="email"
            className="border bg-gray-600 border-gray-400 rounded-md p-2 mt-[8px]"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div className="text-red-500">{error}</div>
        <button
          className="bg-blue-500 text-white cursor-pointer rounded-md p-2 mt-[1em]"
          onClick={save}
        >
          save
        </button>
      </form>
    </main>
  )
}