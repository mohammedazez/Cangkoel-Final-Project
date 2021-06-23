import {
	SET_NAMA,
	SET_NOMOR_HP,
	SET_DOKUMEN_PERIZINAN,
	SET_NOMOR_NPWP,
	SET_KTP,
	SET_JENIS_USAHA,
	SET_TENAGA_KERJA,
	SET_OMSET,
	SET_ALAMAT,
	SET_PETANI_ID
} from './pengajuanActionTypes'

const inititalState = {
	nama: '',
	nomorHP: '',
	dokumenPerizinan: '',
	nomorNPWP: '',
	ktp: '',
	jenisUsaha: '',
	tenagaKerja: '',
	omset: '',
	alamat: '',
	petaniID: ''
}

const userRegisterReducer = (state = inititalState, action) => {
	switch (action.type) {
		case SET_NAMA:
			return {
				...state,
				nama: action.payload.nama
			}
		case SET_NOMOR_HP:
			return {
				...state,
				nomorHP: action.payload.nomorHP
			}
		case SET_DOKUMEN_PERIZINAN:
			return {
				...state,
				dokumenPerizinan: action.payload.dokumenPerizinan
			}
		case SET_NOMOR_NPWP:
			return {
				...state,
				nomorNPWP: action.payload.nomorNPWP
			}
		case SET_KTP:
			return {
				...state,
				ktp: action.payload.ktp
			}
		case SET_JENIS_USAHA:
			return {
				...state,
				jenisUsaha: action.payload.jenisUsaha
			}
		case SET_TENAGA_KERJA:
			return {
				...state,
				tenagaKerja: action.payload.tenagaKerja
			}
		case SET_OMSET:
			return {
				...state,
				omset: action.payload.omset
			}
		case SET_ALAMAT:
			return {
				...state,
				alamat: action.payload.alamat
			}
		case SET_PETANI_ID:
			return {
				...state,
				petaniID: action.payload.petaniID
			}
		default:
			return state
	}
}

export default userRegisterReducer