import useRepository from '@/hooks/useRepository'
import { User } from '@/type/user.type'
import axios from 'axios'
import { useEffect, useState } from 'react'
import DeleteButton from './buttons/DeleteButton'
import EmailEditInput from './inputs/EmaiEditlInput'
import NameEditInput from './inputs/NameEditlInput'
import UpdateButton from './buttons/UpdateButton'

type Props = {
  clear: () => void
  users: { [key: number]: User }
  setUsers: React.Dispatch<React.SetStateAction<{ [key: number]: User }>>
  editUsers: { [key: number]: User }
  setEditUsers: React.Dispatch<React.SetStateAction<{ [key: number]: User }>>
  setError: React.Dispatch<React.SetStateAction<string>>
}

const UserList: React.FC<Props> = ({
  clear,
  users,
  setUsers,
  editUsers,
  setEditUsers,
  setError,
}) => {
  const [editMode, setEditMode] = useState<Map<number, boolean>>(new Map())
  const repository = useRepository()

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const users = await repository.user.findAll()
        const map: { [key: number]: User } = users.reduce(
          (acc, user) => ({
            ...acc,
            [user.id as number]: user,
          }),
          {},
        )
        setUsers(map)
        setEditUsers(map)
      } catch (error) {
        console.error(error)
        if (axios.isAxiosError(error) && error.response) {
          // 404エラー以外はエラー出力する
          if (error.response.status !== 404) {
            console.error(error.message)
            setError(error.message)
          }
        }
      }
    }
    fetchUsers()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  const update = async (id: number) => {
    const user: User = {
      id,
      username: editUsers[id].username || '',
      email: editUsers[id].email || '',
    }
    try {
      await repository.user.update(user)
      clear()
      toggleEditMode(id)
      setUsers((prev) => ({
        ...prev,
        [id]: {
          ...prev[id],
          username: user.username,
          email: user.email,
        },
      }))
    } catch (error) {
      if (axios.isAxiosError(error) && error.response) {
        setError(error.response.data.message)
      }
      console.error(error)
    }
  }

  const userDelete = async (id: number) => {
    try {
      await repository.user.delete(id)
      toggleEditMode(id)
      setUsers((prev) => {
        const newUsers = { ...prev }
        delete newUsers[id]
        return newUsers
      })
    } catch (error) {
      if (axios.isAxiosError(error) && error.response) {
        setError(error.response.data.message)
      }
      console.error(error)
    }
  }

  const toggleEditMode = (id: number) => {
    setEditMode((prev) => {
      const newEditMode = new Map(prev)
      newEditMode.set(id, !newEditMode.get(id))
      return newEditMode
    })
  }

  return (
    <ul>
      {Object.keys(users)
        .map((key) => parseInt(key))
        .sort((a, b) => a - b)
        .map((key) => (
          <li key={key}>
            <pre className="mt-1">
              {editMode.get(users[key].id as number) ? (
                <span>
                  <NameEditInput
                    userId={users[key].id as number}
                    editUsers={editUsers}
                    setEditUsers={setEditUsers}
                  />
                  <EmailEditInput
                    userId={users[key].id as number}
                    editUsers={editUsers}
                    setEditUsers={setEditUsers}
                  />
                </span>
              ) : (
                <>
                  Name: {users[key].username}&#009;&#009;Email:{' '}
                  {users[key].email}
                </>
              )}

              <button
                className="bg-gray-500 rounded-md px-1 py-[2px] ml-[1em]"
                onClick={() => toggleEditMode(users[key].id as number)}
              >
                {editMode.get(users[key].id as number) ? 'Cancel' : 'Edit'}
              </button>
              {editMode.get(users[key].id as number) && (
                <UpdateButton id={users[key].id as number} onClick={update} />
              )}
              {editMode.get(users[key].id as number) ? (
                <></>
              ) : (
                <DeleteButton
                  id={users[key].id as number}
                  onClick={userDelete}
                />
              )}
            </pre>
          </li>
        ))}
    </ul>
  )
}

export default UserList
