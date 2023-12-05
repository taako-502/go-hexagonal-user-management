type Props = {
  email: string
  setEmail: React.Dispatch<React.SetStateAction<string>>
}

const EmailInput: React.FC<Props> = ({ email, setEmail }) => (
  <>
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
  </>
)

export default EmailInput
