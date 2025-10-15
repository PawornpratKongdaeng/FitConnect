import "./styles/globals.css"
import { Routes, Route, Navigate } from 'react-router-dom'
import Home from './pages/Home'
import Trainers from './pages/Trainers'
import GymsNear from './pages/GymsNear'
import Login from './pages/Login'


export default function App() {
return (
<Routes>
<Route path="/" element={<Home />} />
<Route path="/login" element={<Login />} />
<Route path="/trainers" element={<Trainers />} />
<Route path="/gyms/near" element={<GymsNear />} />
<Route path="*" element={<Navigate to="/" />} />
</Routes>
)
}