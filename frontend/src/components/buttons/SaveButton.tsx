type Props = {
  onClick: () => Promise<void>
}

const SaveButton: React.FC<Props> = ({ onClick }) => (
  <button
    className="bg-blue-500 text-white cursor-pointer rounded-md p-2 mt-[1em]"
    onClick={onClick}
  >
    save
  </button>
)

export default SaveButton
