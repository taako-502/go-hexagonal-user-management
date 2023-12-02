import { User } from '@/type/user.type'
import { Dispatch, SetStateAction } from 'react'

interface Props {
  userId: number
  editUsers: { [key: number]: User }
  setEditUsers: Dispatch<SetStateAction<{ [key: number]: User }>>
}

const NameEditInput: React.FC<Props> = ({
  userId,
  editUsers,
  setEditUsers,
}) => {
  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const updatedUser: User = {
      ...editUsers[userId],
      username: e.target.value,
    }

    setEditUsers((prevState) => ({
      ...prevState,
      [userId]: updatedUser,
    }))
  }

  return (
    <>
      <label htmlFor={`edit_username_${userId}`} className="mr-[2em]">
        Name:
      </label>
      <input
        id={`edit_username_${userId}`}
        type="text"
        className="border bg-gray-600 border-gray-400 rounded-md p-2 mt-[8px]"
        value={editUsers[userId]?.username}
        onChange={handleNameChange}
      />
    </>
  )
}

export default NameEditInput
