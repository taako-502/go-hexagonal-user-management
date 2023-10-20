import { User } from '@/type/user.type'
import axios from 'axios'
import { useState } from 'react'

export default function Home() {
  const [user, setUser] = useState('')

  const save = async () => {
    const body: User = {
      Username: 'test',
      Email: 'test@test.com',
    }
    try {
      console.log('process.env.NEXT_PUBLIC_BASE_URL')
      console.log(process.env.NEXT_PUBLIC_BASE_URL)
      axios.defaults.baseURL = process.env.NEXT_PUBLIC_BASE_URL
      await axios.post('/user', body)
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <h1 className="text-2xl">ユーザ管理画面</h1>
      <input
        type="text"
        className="border bg-gray-600 border-gray-400 rounded-md p-2 mt-[1em]"
        value={user}
        onChange={(e) => setUser(e.target.value)}
      />
      <button
        className="bg-blue-500 text-white rounded-md p-2 mt-[1em]"
        onClick={save}
      >
        save
      </button>
    </main>
  )
}
