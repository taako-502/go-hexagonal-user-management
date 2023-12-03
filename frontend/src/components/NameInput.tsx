type Props = {
  username: string
  setUsername: React.Dispatch<React.SetStateAction<string>>
}

const NameInput: React.FC<Props> = ({ username, setUsername }) => (
  <>
    {' '}
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
  </>
)

export default NameInput
