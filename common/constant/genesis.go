// Constants to generate/validate genesis block
// Note: this file has been auto-generated by 'genesis generate' command
package constant

const (
	MainchainGenesisBlockID int64 = 2671113494717183072
)

type (
	GenesisConfigEntry struct {
		AccountAddress     string
		AccountBalance     int64
		NodePublicKey      []byte
		NodeAddress        string
		LockedBalance      int64
		ParticipationScore int64
	}
)

var (
	ApplicationCodeName            = "ZBC_beta"
	ApplicationVersion             = "1.0.0"
	MainchainGenesisBlocksmithID   = []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	MainchainGenesisBlockSignature = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	MainchainGenesisTransactionSignature = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	MainchainGenesisBlockTimestamp = int64(1596708000)
	MainchainGenesisAccountAddress = "ZBC_AQTEH7K4_L45WJPLL_HCEC65ZH_7XC5N3XD_YNKPHK45_POH7PQME_AFAFBDWM"
	MainchainGenesisBlockSeed      = make([]byte, 64)
	MainchainGenesisNodePublicKey  = make([]byte, 32)
	GenesisConfig                  = []GenesisConfigEntry{
		{
			AccountAddress: "ZBC_FE7XA5NC_MICAGCZV_XXLS44NV_S4AVWKU6_OTMUMHVT_K5CACFRJ_PTIOBYU5",
			AccountBalance: 0,
			NodePublicKey: []byte{114, 49, 138, 93, 82, 235, 98, 143, 246, 67, 164, 252, 21, 253, 208, 81, 140, 61, 1,
				158, 106, 172, 18, 137, 216, 91, 21, 65, 150, 36, 215, 131},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_ELKDMSXE_YYG3SIJD_NHG2OBMF_SUHE5PJP_WK4IEWXF_MF77R62B_INYZQG4B",
			AccountBalance: 0,
			NodePublicKey: []byte{86, 27, 60, 242, 236, 221, 112, 69, 132, 249, 115, 207, 8, 7, 219, 44, 75, 44, 156,
				200, 249, 238, 109, 145, 24, 215, 229, 241, 135, 153, 121, 59},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_SRIOWBBD_4GY6OZXQ_RHOAE2AH_4CSUNLNP_CUVBLJ24_ZED7UI32_EITV2AOH",
			AccountBalance: 0,
			NodePublicKey: []byte{127, 45, 254, 52, 68, 221, 41, 174, 223, 141, 102, 144, 0, 250, 112, 31, 116, 102, 1,
				93, 177, 45, 61, 76, 187, 83, 67, 162, 12, 84, 14, 30},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_77BY2XWI_ZTIIQUYC_FIU7VP2C_S5CDSZ6Q_IS2CMZFP_CRZJZIFM_XDGTJF2Y",
			AccountBalance: 0,
			NodePublicKey: []byte{29, 214, 213, 35, 147, 212, 123, 10, 227, 222, 116, 211, 70, 208, 144, 64, 172, 234, 97,
				105, 31, 140, 84, 131, 121, 146, 89, 20, 182, 228, 184, 213},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_4RUA3O5K_62KZQZ26_2NMKWZBW_65XKEE6B_HIOQKTSN_22JQV4GG_3YUETFLP",
			AccountBalance: 0,
			NodePublicKey: []byte{141, 2, 254, 203, 95, 116, 10, 39, 102, 196, 211, 147, 176, 202, 229, 239, 133, 197, 239,
				219, 249, 68, 106, 186, 253, 27, 36, 112, 112, 126, 60, 252},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_JRUZ4DF2_6NZT4IW7_4GBOW22T_Q2EB6ERH_W67NN3VA_XALYYSII_XGHMEL2K",
			AccountBalance: 0,
			NodePublicKey: []byte{114, 64, 30, 84, 204, 15, 240, 107, 101, 106, 59, 192, 227, 57, 83, 61, 143, 151, 63,
				141, 117, 209, 173, 46, 236, 189, 224, 255, 188, 101, 208, 106},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_VXRUSCTL_Y3WR57EM_ZJ2CU6MU_ZL2VK32L_EAWEADBX_MZV7ERQ5_PXKEE6YS",
			AccountBalance: 0,
			NodePublicKey: []byte{50, 191, 27, 133, 38, 167, 29, 63, 240, 32, 129, 100, 95, 42, 202, 77, 28, 131, 49,
				70, 16, 167, 199, 105, 165, 189, 110, 243, 21, 122, 82, 207},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_GM3QVQ7G_K6WGJMG3_S6XHFOIW_73SPMYAH_C4FM2AZK_H55OKRYV_YCR3DWBV",
			AccountBalance: 0,
			NodePublicKey: []byte{39, 77, 211, 67, 141, 87, 125, 66, 222, 241, 135, 205, 211, 54, 88, 113, 173, 43, 147,
				185, 25, 234, 6, 27, 139, 174, 25, 41, 69, 142, 240, 194},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_7VYYXSTX_VWZGJRIW_ZV777HYJ_4T4DTPGX_IBAQHGQP_52VQOJKS_G5DBUBE3",
			AccountBalance: 0,
			NodePublicKey: []byte{134, 139, 211, 205, 22, 89, 8, 234, 13, 83, 10, 95, 251, 109, 0, 101, 140, 253, 35,
				202, 206, 36, 102, 209, 153, 172, 152, 222, 46, 119, 137, 85},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_6QQ27H6K_YV4734VK_R76HGITL_PNMEQFH5_JQFQLPTE_KRKRTDVQ_6CRSH2UF",
			AccountBalance: 0,
			NodePublicKey: []byte{116, 185, 173, 173, 109, 248, 162, 163, 235, 144, 126, 25, 80, 105, 99, 243, 139, 32, 65,
				192, 72, 242, 139, 232, 247, 107, 16, 243, 232, 33, 117, 156},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_EHZGAVSA_L2FVJXZ4_7XKD6TST_FLS4WYCI_L25RAAMB_6MV6NQY5_ZYLQ3DBC",
			AccountBalance: 0,
			NodePublicKey: []byte{223, 53, 246, 205, 4, 206, 25, 106, 87, 83, 81, 24, 203, 168, 130, 15, 185, 26, 146,
				216, 22, 155, 243, 156, 89, 53, 224, 153, 21, 8, 94, 60},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_7HI6ZLRF_AALOBQYC_4CSAYH4E_RXG23F57_GJK6Y3CE_UXJ6PRSH_BRHRVIYE",
			AccountBalance: 0,
			NodePublicKey: []byte{104, 85, 46, 193, 9, 8, 67, 198, 89, 172, 146, 10, 149, 44, 184, 180, 162, 195, 248,
				133, 149, 172, 221, 236, 212, 255, 23, 162, 11, 239, 133, 209},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_YCRZ2324_3L55C76N_B6KGQB5Y_RSGDTQ56_N2AMZ32V_QZK7E7PP_ZPICMJ3B",
			AccountBalance: 0,
			NodePublicKey: []byte{83, 235, 251, 130, 193, 36, 90, 215, 45, 194, 70, 243, 48, 50, 175, 168, 91, 147, 70,
				63, 90, 50, 131, 240, 72, 197, 223, 167, 51, 206, 244, 145},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_AIUFCYKZ_MGVTVL4R_EKJMEDKF_CIVU454K_HEQT6VAD_KPYCUGSA_ZW4VHACC",
			AccountBalance: 0,
			NodePublicKey: []byte{206, 80, 152, 66, 74, 68, 119, 217, 152, 26, 234, 247, 114, 63, 226, 151, 107, 140, 41,
				0, 11, 169, 211, 104, 12, 251, 38, 29, 46, 192, 186, 202},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_K6PETTSM_6WTVZ4CO_5ITI6HDO_A2R2Z5HN_LDNVUUJH_NJVYBSLC_JG5DG7ZZ",
			AccountBalance: 0,
			NodePublicKey: []byte{241, 47, 85, 44, 244, 99, 217, 236, 150, 56, 179, 228, 96, 149, 11, 150, 83, 25, 254,
				140, 70, 227, 177, 66, 181, 39, 182, 52, 146, 176, 170, 222},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_66I3SPOK_C7TJW32O_PHZHFCPK_OASFOTKF_2OFJEWTX_G6YBXIA7_4NDUGFJF",
			AccountBalance: 0,
			NodePublicKey: []byte{47, 118, 178, 21, 182, 172, 126, 121, 255, 48, 13, 52, 28, 216, 176, 146, 48, 148, 135,
				231, 140, 41, 144, 194, 68, 84, 94, 204, 55, 88, 51, 158},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_II5COME5_K2VTR54C_UWINLNKL_3ERZ2APX_P6LZFKIU_QDH5YEQ3_HIHJO27J",
			AccountBalance: 0,
			NodePublicKey: []byte{56, 69, 144, 237, 105, 129, 111, 55, 14, 1, 92, 94, 96, 103, 81, 127, 60, 79, 34,
				85, 152, 131, 66, 145, 12, 255, 26, 124, 178, 184, 248, 228},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_JIWBXM3T_Q2XXPBG5_QVHWYCW5_3NRQDKAK_RNWGRT2S_MS46QXU5_RDCBJTBH",
			AccountBalance: 0,
			NodePublicKey: []byte{232, 36, 252, 232, 7, 91, 43, 81, 140, 178, 126, 118, 2, 214, 212, 110, 103, 239, 88,
				215, 142, 185, 147, 248, 100, 250, 143, 138, 55, 79, 1, 47},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_6MUFG4OK_2242C4R6_KLOBYKXC_7QNF6REK_CTEFNT2U_ENLVSE5Y_Y235RYCU",
			AccountBalance: 0,
			NodePublicKey: []byte{25, 23, 48, 26, 80, 90, 1, 189, 130, 41, 117, 58, 233, 219, 6, 122, 45, 32, 89,
				64, 65, 163, 105, 49, 194, 101, 192, 4, 91, 109, 217, 214},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_XI5TNSGA_S4XQOHQI_G3WCNURA_VJ4VIU3F_AHGIUHFG_Z6BDUODD_SP5QYOUY",
			AccountBalance: 0,
			NodePublicKey: []byte{214, 236, 200, 195, 42, 62, 19, 52, 67, 252, 207, 2, 124, 201, 95, 8, 98, 234, 123,
				232, 192, 75, 90, 25, 158, 249, 177, 71, 151, 28, 217, 50},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_WQJ5O2MU_CRHVZDKC_OJI7LHJO_YIOJDZBU_IW5ZR2TT_O3BXCYEC_XGFEFIKB",
			AccountBalance: 0,
			NodePublicKey: []byte{207, 219, 148, 172, 206, 17, 236, 97, 249, 85, 157, 24, 55, 205, 184, 39, 204, 36, 41,
				96, 178, 53, 163, 247, 139, 206, 136, 35, 65, 66, 225, 107},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_XHNZBP7F_B2W2KQAK_JT4XCX4W_QISTAWV6_Q5GZDIG7_SIAQPGWY_APPP2U5D",
			AccountBalance: 0,
			NodePublicKey: []byte{249, 233, 39, 71, 218, 190, 160, 57, 56, 92, 183, 182, 141, 169, 51, 4, 13, 224, 102,
				100, 224, 44, 244, 31, 93, 75, 163, 237, 74, 186, 154, 245},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_W25G72A3_3PBKZ2GQ_MNAAWDG2_P2GURQN3_3F757OAU_CCLUXYLE_6STAXMRM",
			AccountBalance: 0,
			NodePublicKey: []byte{90, 238, 127, 150, 104, 144, 179, 244, 221, 203, 154, 138, 173, 66, 52, 223, 48, 161, 147,
				53, 202, 183, 21, 236, 153, 90, 99, 40, 127, 114, 183, 210},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_FYX3ADCD_YXFHHNXA_SLDRLOQN_UYYNPNZ6_MOJXTL4N_O7SHNQYO_LHUBGUS4",
			AccountBalance: 0,
			NodePublicKey: []byte{221, 114, 117, 103, 131, 128, 76, 170, 113, 72, 240, 18, 24, 165, 157, 204, 75, 182, 220,
				170, 157, 248, 245, 129, 170, 152, 153, 216, 79, 219, 199, 253},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_ZMOQQWQP_3A5GKQCI_Y4ZYWL32_P7EHXS3O_P742SBTQ_BVFCKAN5_CIO2NKCF",
			AccountBalance: 0,
			NodePublicKey: []byte{49, 96, 33, 246, 190, 17, 7, 181, 218, 241, 120, 237, 168, 239, 129, 220, 239, 207, 162,
				25, 210, 13, 100, 194, 38, 17, 2, 33, 212, 117, 234, 147},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_SQNDE2HE_WLFUY6AS_BE46WZED_Q7AYDDWS_K3G3H37Z_COQ53PS2_LMGA6YF4",
			AccountBalance: 0,
			NodePublicKey: []byte{34, 81, 42, 35, 199, 103, 96, 79, 194, 155, 189, 175, 83, 202, 91, 246, 128, 152, 102,
				136, 224, 161, 124, 3, 16, 189, 59, 197, 174, 114, 239, 25},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_HKT25PQI_7PVGDNXZ_YIXRITGZ_IHEJGIRC_GTPYWE3U_6HXTVZRX_GAQY4QKO",
			AccountBalance: 0,
			NodePublicKey: []byte{66, 216, 81, 215, 63, 62, 75, 8, 12, 104, 38, 16, 235, 122, 23, 45, 0, 205, 66,
				82, 222, 204, 160, 25, 39, 190, 251, 148, 150, 156, 232, 106},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_IOOHGCTM_XO3MLEMR_FTRZQ52X_CXTUMELH_5LMKJVUO_TMSUOLIO_FQGABWNP",
			AccountBalance: 0,
			NodePublicKey: []byte{213, 53, 227, 21, 178, 100, 172, 146, 113, 31, 189, 217, 247, 36, 37, 218, 253, 51, 174,
				129, 132, 106, 85, 210, 135, 152, 40, 252, 200, 154, 82, 129},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_SOLGFQ5L_FLBT5S7T_YS7BE76Z_NJXTYQ6O_DTDIHF3K_OX2LSZCW_ZLUF4PDY",
			AccountBalance: 0,
			NodePublicKey: []byte{49, 60, 25, 83, 74, 47, 85, 10, 12, 167, 38, 70, 75, 92, 191, 53, 32, 68, 165,
				45, 92, 142, 133, 178, 79, 217, 60, 74, 45, 206, 38, 234},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_NA4LB3KJ_QQMHYKHO_557HUJWZ_HPQUNDN2_KUYHA2VX_D5SRCKGX_WWL674X7",
			AccountBalance: 0,
			NodePublicKey: []byte{198, 146, 12, 24, 64, 77, 162, 255, 13, 91, 107, 0, 178, 120, 111, 141, 31, 202, 108,
				227, 119, 88, 144, 212, 123, 73, 21, 204, 103, 141, 245, 253},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_HAFL5PZB_QT7OKLTI_FJ4DR7AZ_7KZKSJKS_TM5NAHCW_R4MI6NY7_6HDQI7HI",
			AccountBalance: 0,
			NodePublicKey: []byte{10, 175, 110, 255, 122, 220, 3, 67, 78, 5, 105, 245, 210, 236, 42, 223, 173, 66, 182,
				193, 213, 173, 148, 171, 0, 3, 151, 254, 148, 73, 182, 60},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_QOCURDZV_XBMTWBLO_MWHL7DNB_Y3AJXO6X_OMJ6FCJ2_C2XOKUVK_KGUFVMI7",
			AccountBalance: 0,
			NodePublicKey: []byte{247, 5, 158, 252, 227, 143, 89, 179, 97, 181, 197, 67, 201, 184, 235, 147, 223, 251, 190,
				110, 179, 59, 1, 33, 116, 60, 150, 76, 208, 244, 8, 193},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_F4TLGXGT_KEFT4IVU_3LQ77K64_MD7NLJ7E_RBT5JDL6_26LA6SDT_Z2XFJZYB",
			AccountBalance: 0,
			NodePublicKey: []byte{53, 217, 11, 212, 20, 234, 204, 234, 29, 57, 216, 201, 108, 55, 18, 195, 19, 224, 82,
				125, 252, 31, 87, 153, 159, 207, 62, 143, 84, 246, 160, 124},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_NH6TNI2W_HR2LRNUQ_5EM7MJQ2_5A4AIF7R_TOSDY5XF_3UHITAWQ_TWY2CZTA",
			AccountBalance: 0,
			NodePublicKey: []byte{127, 56, 184, 151, 39, 133, 252, 188, 25, 3, 103, 125, 29, 145, 228, 137, 210, 146, 120,
				126, 4, 40, 181, 137, 153, 45, 115, 20, 254, 238, 53, 223},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_7TWMHPDP_AGDG62AH_AGXQVTKN_7WHUPFOX_4UNMYSNJ_WNORVKVJ_464OFJHY",
			AccountBalance: 0,
			NodePublicKey: []byte{241, 243, 243, 107, 73, 149, 65, 157, 119, 253, 214, 119, 180, 231, 44, 232, 214, 226, 88,
				212, 157, 6, 61, 76, 211, 189, 224, 238, 136, 209, 30, 218},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_QYLUE2X2_W4V7YT7J_FEA3J7RJ_SYWR2FSP_YK7JE3TR_W6UZV5SG_AYUTDCVO",
			AccountBalance: 0,
			NodePublicKey: []byte{95, 76, 51, 117, 166, 91, 186, 10, 5, 87, 106, 62, 159, 112, 177, 214, 74, 72, 220,
				132, 0, 194, 232, 181, 50, 208, 169, 59, 15, 174, 9, 125},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_FONKWTQQ_MEDSEN3D_ULLB4MIG_UHEEUTJR_ZVX22Y2X_BYYA44NZ_MABIOTYQ",
			AccountBalance: 0,
			NodePublicKey: []byte{233, 224, 235, 84, 75, 113, 236, 51, 186, 57, 60, 189, 67, 2, 6, 2, 5, 127, 111,
				27, 92, 35, 155, 80, 149, 201, 212, 229, 155, 245, 122, 42},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_3SYU5NST_B37UDLCA_QEFKBDTM_U5BTGSRB_R7UNPDSD_2HJ5DHR2_MZJDXGPR",
			AccountBalance: 0,
			NodePublicKey: []byte{85, 21, 217, 205, 171, 156, 41, 156, 226, 98, 192, 142, 193, 242, 78, 73, 213, 241, 123,
				168, 98, 82, 109, 117, 188, 56, 24, 203, 67, 240, 142, 34},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_OZNFCB26_JJOBFVH5_KJF7SS5W_TPWKPX5R_VMO5AUXL_GH6BW2DW_2S3DAG27",
			AccountBalance: 0,
			NodePublicKey: []byte{139, 52, 108, 133, 13, 12, 73, 127, 164, 34, 205, 23, 221, 124, 173, 95, 66, 89, 239,
				125, 109, 8, 66, 163, 204, 15, 119, 237, 121, 226, 230, 145},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_BZZS2OIT_AHKDL6BM_G66WGTZP_SJUAET5D_Z675WLCG_ZBAKFFUE_5KWOZ7VS",
			AccountBalance: 0,
			NodePublicKey: []byte{32, 144, 26, 243, 187, 215, 140, 32, 115, 47, 230, 173, 138, 149, 84, 86, 186, 8, 86,
				75, 139, 6, 205, 9, 19, 10, 53, 109, 53, 208, 115, 223},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_FBQ6FMCG_7VGDPCDS_ZKZNQ73V_M2DTAJP5_TWXCQGBR_7MA3GPB4_NNEZS3LR",
			AccountBalance: 0,
			NodePublicKey: []byte{15, 41, 198, 182, 167, 117, 5, 223, 6, 60, 249, 207, 19, 107, 185, 228, 51, 166, 61,
				187, 177, 47, 105, 165, 75, 238, 103, 255, 120, 106, 161, 102},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_GJRJZ5BJ_CVT3CVXW_JJYEDNCJ_K2AZKGQ3_QEYA7DUA_RPNEX7RH_3W4DU3FE",
			AccountBalance: 0,
			NodePublicKey: []byte{111, 6, 172, 107, 152, 234, 12, 246, 221, 183, 22, 128, 208, 252, 226, 191, 66, 174, 47,
				255, 223, 245, 221, 129, 113, 224, 100, 59, 1, 33, 87, 243},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_RZAMPLER_JSWUPX3Z_5C45IHAA_ZLSBY22Z_RPMYHU5I_UHVK7IYA_OW4UKTJS",
			AccountBalance: 0,
			NodePublicKey: []byte{213, 187, 5, 144, 196, 42, 164, 44, 172, 141, 101, 30, 215, 171, 72, 191, 14, 122, 179,
				224, 47, 68, 0, 55, 216, 227, 48, 154, 239, 85, 217, 81},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_6S667AWI_VDVZ7MZM_TGZNICF7_55DWHYWU_AGJ2SGUS_AOKOVJSB_VETEY45C",
			AccountBalance: 0,
			NodePublicKey: []byte{154, 47, 39, 73, 218, 24, 97, 219, 234, 234, 17, 187, 15, 198, 210, 117, 79, 72, 21,
				157, 185, 226, 26, 249, 87, 217, 167, 199, 128, 37, 112, 97},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_IM3OTLLJ_AOU7MYDO_VWBHW7SG_4ZWVCJMP_4PF6JLUG_D5TK57IJ_DHBUYIVA",
			AccountBalance: 0,
			NodePublicKey: []byte{204, 76, 235, 204, 185, 211, 9, 8, 8, 105, 254, 80, 59, 23, 234, 160, 244, 166, 29,
				47, 182, 239, 236, 236, 100, 148, 243, 79, 161, 178, 143, 243},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_FMTJNDJE_6NX2N5HP_CFC7OF2L_X6QNNWU7_LJNE6U4Q_PPCQ5ZDS_2XRSC33V",
			AccountBalance: 0,
			NodePublicKey: []byte{174, 10, 11, 111, 242, 254, 58, 161, 246, 86, 25, 238, 147, 153, 251, 220, 74, 82, 49,
				22, 50, 143, 48, 160, 42, 178, 166, 93, 179, 163, 238, 59},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_HWDLORS7_ANOMNO2Y_LMK3KZSO_3AG2FF4H_YO5HCI6S_3OFKIMUG_RLZHJXZH",
			AccountBalance: 0,
			NodePublicKey: []byte{149, 66, 123, 154, 120, 144, 78, 4, 40, 236, 116, 50, 163, 3, 179, 146, 163, 248, 148,
				111, 230, 147, 237, 230, 74, 85, 212, 112, 92, 155, 202, 86},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_XNOBFSH5_HUYO7GA6_HAXJMYGZ_VDGBCBYA_6XPGNLCH_NAWOKPGY_O5RTY4FE",
			AccountBalance: 0,
			NodePublicKey: []byte{67, 70, 187, 159, 112, 188, 146, 35, 80, 139, 193, 151, 6, 115, 208, 147, 217, 79, 148,
				195, 164, 203, 129, 227, 93, 67, 95, 253, 195, 97, 182, 68},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_ND55K2ZK_TO7ZBZBT_7NFEAUEJ_FTYWVFHH_D7MYWQKD_XR67LKAY_NXF6IL7Y",
			AccountBalance: 0,
			NodePublicKey: []byte{183, 33, 123, 68, 153, 105, 79, 163, 195, 10, 57, 113, 244, 0, 83, 85, 15, 163, 146,
				186, 51, 147, 136, 131, 150, 211, 53, 240, 37, 145, 47, 222},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_TKBQGGEA_S2VUQSP5_ELVLRMHB_2ECIZDBP_WEYVTCML_JFBHLNHS_XIBW77HC",
			AccountBalance: 0,
			NodePublicKey: []byte{136, 191, 73, 196, 238, 129, 10, 103, 113, 73, 231, 122, 90, 87, 17, 31, 177, 104, 117,
				225, 75, 83, 78, 251, 40, 179, 220, 214, 134, 24, 127, 18},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_HQVY2UYX_IRMMZXLT_6LL2O7DM_G5D62FCQ_IK5F6PYX_VCFH5E3M_GBPHKLE6",
			AccountBalance: 0,
			NodePublicKey: []byte{203, 100, 7, 91, 123, 51, 37, 26, 3, 184, 17, 100, 55, 110, 181, 163, 131, 186, 218,
				89, 67, 35, 6, 98, 11, 65, 120, 23, 222, 78, 112, 105},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_PUBYABEN_J6LE5S6M_I6BALJ5W_WLPI4VLW_EYPZR6AH_4ORX76OO_VSMYSW57",
			AccountBalance: 0,
			NodePublicKey: []byte{89, 69, 26, 11, 141, 183, 90, 16, 24, 218, 195, 227, 235, 14, 221, 3, 77, 132, 151,
				98, 178, 212, 59, 185, 170, 218, 121, 103, 21, 29, 54, 250},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_UJCZ4BGN_HT2CY5DH_BQ5B7VQC_RP5YASYA_2M73WKYK_ZX5CTOU7_KCMIYTJP",
			AccountBalance: 0,
			NodePublicKey: []byte{128, 170, 8, 199, 9, 44, 122, 83, 194, 59, 194, 18, 237, 22, 73, 29, 238, 181, 209,
				232, 88, 196, 214, 167, 219, 83, 144, 35, 188, 5, 218, 221},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_ALJ65FI5_MW4J7J5H_2YOC6N5K_3CCCUQGA_57YMWZDS_GIJMJFVB_K5VY4T5D",
			AccountBalance: 0,
			NodePublicKey: []byte{230, 210, 116, 27, 151, 85, 126, 208, 39, 181, 22, 243, 108, 76, 180, 54, 237, 76, 72,
				154, 40, 201, 143, 121, 126, 233, 150, 252, 156, 196, 173, 119},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_QPRTQAEH_MNMBLV6H_GIPLBJ6O_ZVUN2JAV_VNRREKOQ_4B262SIS_SDGFYWQT",
			AccountBalance: 0,
			NodePublicKey: []byte{196, 77, 19, 231, 241, 255, 135, 250, 99, 164, 190, 196, 98, 209, 53, 130, 44, 250, 76,
				223, 0, 30, 225, 7, 111, 217, 182, 100, 212, 202, 255, 184},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_LKZTP3D5_IF7NIB6C_47F5QLMH_TOMYO246_BBBYYQAO_5M6UKLEL_2HC5RP2G",
			AccountBalance: 0,
			NodePublicKey: []byte{23, 150, 215, 135, 240, 151, 252, 63, 68, 7, 29, 166, 121, 150, 17, 28, 49, 66, 239,
				8, 56, 170, 93, 14, 34, 131, 205, 29, 96, 236, 187, 186},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_YNXLCUYD_KBLDIVHH_PVT6EMQE_6HZEIRFG_US2JA7QP_MKTI7NQX_IZGBCFGS",
			AccountBalance: 0,
			NodePublicKey: []byte{42, 252, 199, 190, 181, 204, 10, 125, 153, 7, 136, 253, 227, 55, 1, 127, 119, 165, 5,
				33, 34, 184, 82, 90, 89, 185, 29, 159, 142, 150, 122, 191},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_VD5FB5TM_ANJXGA3U_VIF3JJXN_M4GAWMKD_Y442LNI7_GMDDVBNY_HIDVHE7B",
			AccountBalance: 0,
			NodePublicKey: []byte{162, 184, 190, 104, 143, 161, 127, 60, 83, 254, 71, 212, 123, 185, 245, 155, 224, 150, 103,
				120, 168, 40, 226, 188, 239, 116, 137, 48, 57, 244, 68, 116},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_FEP2LTNF_FE6QO2D4_FL7QBRJ4_A3H74XM3_H6SJHOXV_2ET6RCEY_JU4X2AFW",
			AccountBalance: 0,
			NodePublicKey: []byte{13, 50, 18, 105, 218, 192, 147, 109, 180, 86, 47, 185, 63, 11, 216, 236, 189, 15, 212,
				1, 19, 10, 26, 92, 52, 203, 56, 57, 124, 55, 13, 93},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_3X5QUQSS_5RXTDUMP_2UWW3ORU_WZPAZBDB_ABTG3S5Z_WCHSO4FB_ABEX7KPL",
			AccountBalance: 0,
			NodePublicKey: []byte{101, 212, 29, 11, 223, 150, 33, 92, 151, 3, 156, 227, 221, 239, 28, 140, 253, 2, 114,
				29, 213, 43, 157, 123, 202, 210, 173, 6, 141, 97, 93, 253},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_IVBG2XF6_OYOCF3N7_HX6WPGMW_7VXSBPGG_OSSKDE6K_KNRLCRA4_NE7G2COR",
			AccountBalance: 0,
			NodePublicKey: []byte{182, 65, 147, 37, 187, 158, 15, 204, 129, 14, 71, 127, 147, 1, 236, 54, 147, 88, 110,
				44, 238, 32, 223, 197, 133, 97, 29, 170, 52, 13, 33, 213},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_2IT7JHWI_BOG7J6TR_ULT7AEZY_UE4QDJWH_PZ65GN7E_A3ASAEEJ_4LTSPR7W",
			AccountBalance: 0,
			NodePublicKey: []byte{208, 26, 170, 17, 128, 205, 52, 35, 66, 218, 173, 172, 235, 161, 141, 107, 13, 177, 125,
				64, 19, 103, 13, 181, 106, 127, 148, 223, 215, 88, 10, 81},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_YE6AWFWQ_OAMFDGSM_X6T5VCOT_5VOC7QZS_WE2CWRRE_FBCW4LMC_JCX2HYIL",
			AccountBalance: 0,
			NodePublicKey: []byte{11, 104, 168, 68, 7, 237, 108, 234, 173, 199, 95, 149, 109, 120, 178, 104, 131, 133, 94,
				214, 76, 243, 95, 61, 123, 168, 161, 198, 173, 171, 33, 5},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_F4DNR6XM_YGI6UNGR_ZOOWSZ6Q_IKMTK4RK_XKXTTKW2_YWOWMACA_FKZIAQKF",
			AccountBalance: 0,
			NodePublicKey: []byte{207, 78, 124, 120, 130, 160, 112, 210, 246, 174, 108, 184, 62, 205, 106, 109, 118, 20, 180,
				95, 234, 229, 101, 110, 203, 101, 219, 38, 188, 80, 178, 192},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_VBSZ26CB_BIXQ6NWJ_FRHQX33T_TZEZX4YN_N3GETGDQ_DODOZWM2_DB6FMVWH",
			AccountBalance: 0,
			NodePublicKey: []byte{214, 76, 12, 156, 38, 46, 22, 20, 65, 210, 239, 209, 191, 123, 235, 128, 37, 186, 63,
				114, 250, 200, 109, 160, 234, 2, 237, 197, 232, 204, 110, 200},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_CXNA4ICA_5JN7ONUD_QCI5UBID_RREOVXPB_SQ7TSLGR_NRNI2UXL_QMQ7L4EY",
			AccountBalance: 0,
			NodePublicKey: []byte{165, 104, 68, 245, 112, 244, 115, 172, 54, 21, 123, 117, 215, 82, 102, 150, 185, 25, 109,
				43, 140, 241, 158, 19, 170, 89, 147, 202, 182, 252, 190, 203},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_6FX33STR_3XZOR353_2XI45KXE_RXEX7CZC_6TITMRFY_7C3VHSMI_RR5GW2DJ",
			AccountBalance: 0,
			NodePublicKey: []byte{198, 55, 249, 78, 222, 161, 34, 196, 163, 13, 101, 159, 24, 35, 167, 157, 153, 124, 225,
				126, 43, 238, 32, 198, 245, 219, 147, 122, 141, 81, 112, 129},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_WJVYDE7Z_OHZM3T66_EI4ZOJTG_MF35RPJ5_EOMBQ2U4_TYOQPDYK_2MMNIY3O",
			AccountBalance: 0,
			NodePublicKey: []byte{13, 30, 182, 167, 131, 99, 0, 176, 180, 5, 162, 120, 206, 101, 226, 167, 67, 114, 141,
				33, 73, 42, 2, 31, 224, 163, 231, 21, 8, 155, 73, 27},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_NJRYUIHS_YPGINNIS_SSVP4S6J_HVVJSMHQ_QLZ32V6D_ZYTAQHW6_NBWOBMNP",
			AccountBalance: 0,
			NodePublicKey: []byte{101, 192, 101, 36, 215, 23, 217, 58, 172, 232, 156, 179, 239, 159, 14, 17, 152, 250, 92,
				176, 24, 153, 95, 24, 228, 77, 107, 178, 45, 148, 204, 9},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_FVVH2G4Z_THWNBMGG_I5FRBIBW_NYFRN6VF_UVARAOME_JU7Y6CQW_WMHHKZTH",
			AccountBalance: 0,
			NodePublicKey: []byte{132, 141, 54, 103, 87, 221, 137, 11, 223, 234, 155, 151, 240, 201, 235, 157, 79, 230, 248,
				227, 172, 26, 36, 120, 65, 239, 253, 101, 69, 71, 15, 109},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_WM2N3S7W_YLC4SD2D_EOSJECLV_3OMSHSFO_7TUN3GPG_4JOA6EHV_Y4TYRHEZ",
			AccountBalance: 0,
			NodePublicKey: []byte{53, 140, 225, 87, 149, 84, 149, 14, 50, 176, 173, 73, 102, 188, 5, 1, 2, 153, 136,
				166, 108, 113, 31, 255, 215, 7, 170, 31, 176, 234, 38, 206},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_H3FEF7WA_SGLUT4MJ_7UJM7KJQ_TVMRIS66_O5SRZDQA_HMP6SCSW_63LUKZJ3",
			AccountBalance: 0,
			NodePublicKey: []byte{157, 195, 38, 130, 205, 54, 218, 190, 21, 55, 78, 162, 229, 214, 77, 121, 86, 19, 21,
				87, 177, 216, 183, 230, 125, 156, 107, 223, 158, 72, 140, 81},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_EAEZLD77_N7FHSQDC_CSZAQQR7_US4UOP3K_6P2PKW7W_YXCVHHJT_PVSUMYFL",
			AccountBalance: 0,
			NodePublicKey: []byte{193, 133, 156, 121, 79, 173, 16, 25, 50, 154, 184, 148, 4, 0, 73, 89, 180, 160, 142,
				12, 128, 218, 136, 46, 104, 182, 217, 46, 119, 74, 92, 36},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_VKIYBRYD_X3PJ7V44_YXG7BXP4_5G7AGOH3_NAEGUB3D_HCXMEUAQ_T5YBZBX7",
			AccountBalance: 0,
			NodePublicKey: []byte{41, 106, 225, 62, 105, 178, 117, 91, 18, 52, 6, 99, 78, 46, 62, 111, 238, 70, 165,
				2, 4, 230, 130, 168, 153, 9, 206, 182, 34, 152, 230, 246},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_LBAV3J2H_B5VDB7CR_BVWMGIFW_3OWJAYH6_GYRDTVSC_GQC63AS4_356VPIYK",
			AccountBalance: 0,
			NodePublicKey: []byte{221, 62, 251, 75, 88, 210, 245, 219, 44, 84, 112, 74, 16, 40, 77, 23, 76, 202, 112,
				157, 231, 8, 165, 195, 70, 35, 246, 228, 20, 10, 186, 152},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_MKJUFPHX_L2N7QJE5_URNJ4KRO_YXQJ4QFY_NYUJBOVV_6TJGRRHE_KFQ4PUPA",
			AccountBalance: 0,
			NodePublicKey: []byte{115, 80, 179, 182, 27, 100, 84, 54, 15, 41, 249, 163, 255, 236, 113, 55, 222, 38, 60,
				178, 156, 214, 112, 19, 225, 10, 153, 221, 240, 142, 226, 11},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_PX4AJAPD_FZ7E3LKT_VZCFJL6N_JFXME3A7_EJDPYU6J_OHENWCUQ_MCQGDIJO",
			AccountBalance: 0,
			NodePublicKey: []byte{62, 0, 31, 98, 98, 246, 180, 75, 66, 51, 239, 132, 2, 159, 76, 27, 184, 0, 19,
				2, 15, 46, 35, 43, 134, 37, 25, 157, 115, 209, 196, 169},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_R3HD2QPG_YSTXYDWH_7SEAB65V_KICGUYNV_FXL6FQ5O_MVHGWJJZ_RV6XQYN3",
			AccountBalance: 0,
			NodePublicKey: []byte{4, 140, 81, 134, 242, 185, 204, 129, 123, 221, 191, 30, 228, 171, 50, 79, 88, 238, 243,
				173, 38, 119, 225, 188, 128, 0, 169, 66, 163, 171, 108, 2},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_2IS6CRUI_T3Z7GSVT_NF3Y34XY_VJDFQW67_3HU7MCTG_43ZI2LR6_2IT7FZNV",
			AccountBalance: 0,
			NodePublicKey: []byte{7, 176, 235, 141, 99, 42, 65, 152, 63, 242, 217, 186, 222, 136, 158, 146, 204, 103, 232,
				250, 215, 172, 122, 152, 178, 152, 128, 122, 113, 60, 34, 122},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_4FI4M2BR_3CBXP57V_YP5WILBB_VVSFUEAL_NJEJJC7K_XNZP2DNM_CKK7JDAL",
			AccountBalance: 0,
			NodePublicKey: []byte{82, 150, 46, 83, 96, 13, 30, 149, 148, 140, 2, 212, 230, 87, 29, 207, 172, 7, 69,
				23, 184, 45, 122, 42, 178, 25, 21, 194, 245, 42, 15, 158},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_QUKNQVKX_HPF224YJ_6A3YL4DY_FT4E7XFW_WGTB5SZQ_5Z5V6PF7_RA5WPGKH",
			AccountBalance: 0,
			NodePublicKey: []byte{40, 241, 85, 156, 225, 96, 202, 135, 162, 166, 54, 181, 36, 39, 102, 143, 102, 106, 231,
				209, 240, 112, 225, 74, 234, 208, 9, 60, 75, 255, 96, 148},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_DTJZWD5X_2EZDC5NR_7XKNPY44_AJPIBWHY_MG33CSTK_ZNE3UXIX_R26ITVMW",
			AccountBalance: 0,
			NodePublicKey: []byte{129, 178, 31, 93, 228, 179, 216, 44, 225, 80, 6, 97, 5, 180, 187, 247, 44, 80, 130,
				52, 16, 20, 57, 119, 146, 121, 163, 60, 132, 136, 124, 225},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_7WCIEGAP_NDRPVNHQ_55QWVLNM_NVEZ4L6M_XGECZFNC_BMXGZC44_3JEI2ZUD",
			AccountBalance: 0,
			NodePublicKey: []byte{105, 87, 30, 22, 236, 23, 152, 210, 218, 58, 201, 148, 182, 172, 221, 30, 62, 234, 78,
				3, 128, 126, 135, 143, 172, 77, 142, 200, 242, 28, 126, 50},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_PUSY3LKA_K7PLCOLS_JUGUZSQG_CG2HXXI6_4JLJQAAJ_Q4DPOUHY_PMQGIXKB",
			AccountBalance: 0,
			NodePublicKey: []byte{52, 47, 107, 243, 123, 54, 5, 177, 183, 111, 181, 140, 45, 219, 163, 223, 73, 208, 169,
				231, 131, 184, 19, 126, 71, 66, 223, 44, 245, 152, 16, 214},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_MPJFAZZT_WDH633ED_5ZI5NJRF_HA3I3E6Z_ZNTZAUOO_JYCQNSUE_QRO4ICPX",
			AccountBalance: 0,
			NodePublicKey: []byte{66, 103, 226, 116, 203, 162, 74, 201, 148, 229, 39, 203, 13, 207, 75, 161, 225, 239, 62,
				7, 37, 48, 69, 80, 139, 19, 181, 51, 51, 143, 128, 18},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_7E5YCF2C_B2IWPUCP_LUG3ZFK4_2YGYZ4TH_W3RRWHKK_OHWLOSOH_E7IJTBOT",
			AccountBalance: 0,
			NodePublicKey: []byte{144, 108, 110, 73, 219, 143, 234, 227, 48, 251, 58, 136, 81, 199, 246, 80, 134, 124, 55,
				105, 202, 197, 22, 158, 34, 0, 250, 171, 135, 232, 136, 132},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_K7FA23QT_CKWRX36G_5V7MC5I7_6SQ76ZMZ_U5HMKQAP_ZMLWKLBT_TSOL4OIX",
			AccountBalance: 0,
			NodePublicKey: []byte{89, 248, 224, 24, 189, 237, 221, 25, 54, 102, 216, 193, 227, 70, 187, 173, 149, 21, 123,
				237, 192, 145, 153, 199, 134, 208, 113, 132, 171, 255, 189, 45},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_2HXBPK6H_WMH6EC2K_PZ4H5C74_DUOONG5V_HCB4P3CZ_GFDB4HWU_BS3YKHHY",
			AccountBalance: 0,
			NodePublicKey: []byte{121, 213, 107, 250, 192, 57, 198, 117, 146, 164, 168, 100, 54, 144, 159, 249, 166, 104, 149,
				115, 191, 56, 173, 96, 95, 149, 248, 238, 187, 58, 28, 250},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_YRJJKSNH_7JOSTQLO_6M2YAZH5_ZND67OHY_VNXIIHCS_NUX5BVQW_SHVY3JS7",
			AccountBalance: 0,
			NodePublicKey: []byte{235, 54, 188, 131, 71, 26, 175, 175, 16, 36, 229, 75, 31, 9, 243, 173, 214, 60, 35,
				243, 28, 99, 85, 137, 253, 210, 171, 8, 101, 197, 104, 60},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_IQPLDDPD_XI37L7P5_PCOIKPLN_HOMG6I27_PFVRARFS_FOTUODSG_NMRIJ2RK",
			AccountBalance: 0,
			NodePublicKey: []byte{109, 106, 49, 52, 45, 205, 109, 228, 246, 16, 169, 39, 67, 62, 10, 162, 104, 21, 24,
				225, 85, 0, 89, 168, 132, 16, 192, 240, 93, 23, 205, 22},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_X4CSOQ7F_ROBFP5E3_VISRJH4U_YDTCJEKC_S55U37US_MSBEKJS4_PBVV6NEP",
			AccountBalance: 0,
			NodePublicKey: []byte{251, 50, 13, 149, 202, 235, 94, 114, 208, 245, 93, 28, 213, 72, 58, 20, 198, 96, 45,
				94, 82, 173, 139, 83, 240, 137, 161, 51, 254, 25, 250, 42},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_QNBXZJWE_ON3RDZVT_X2BRXJMX_3WHBAVMH_HLMII42Z_AK37L446_UXWDQREV",
			AccountBalance: 0,
			NodePublicKey: []byte{184, 37, 16, 234, 187, 204, 80, 118, 233, 200, 225, 43, 219, 159, 50, 173, 231, 227, 218,
				67, 113, 52, 156, 43, 42, 198, 79, 171, 62, 11, 106, 216},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_G7GMIXH4_5S7GWK4H_S2RVC7UD_CZX7C4A2_FKRUY5BF_BRZBHNWD_SMRY5COH",
			AccountBalance: 0,
			NodePublicKey: []byte{206, 232, 140, 99, 55, 134, 224, 231, 44, 165, 17, 93, 42, 83, 27, 182, 204, 141, 210,
				103, 201, 160, 21, 38, 16, 128, 201, 103, 36, 62, 210, 213},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_Z74QFIFU_KV26XJCB_5B7QL5KO_DO5YLQTP_GYLXPDL7_G36UOSSO_I722Z5KH",
			AccountBalance: 0,
			NodePublicKey: []byte{9, 91, 75, 32, 27, 69, 30, 71, 135, 25, 37, 214, 194, 157, 207, 125, 0, 169, 204,
				8, 192, 143, 202, 185, 206, 176, 7, 226, 141, 81, 52, 240},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_SLFSD3SJ_F2F2EAB3_5MJPBGQ2_5YJ7K2UM_E3LXKPNG_RP3H2NZT_XHIQL4DC",
			AccountBalance: 0,
			NodePublicKey: []byte{40, 127, 82, 78, 40, 134, 105, 1, 117, 105, 170, 132, 105, 86, 225, 251, 236, 110, 96,
				13, 142, 91, 118, 222, 83, 3, 82, 33, 177, 34, 21, 55},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_WF5KQ2HS_RZZCORDT_NQQN75QO_LO4YLRGK_P65TVSJP_VZNE4AHH_VJBRYE3Q",
			AccountBalance: 0,
			NodePublicKey: []byte{39, 174, 101, 149, 39, 151, 237, 98, 24, 204, 50, 127, 175, 184, 247, 221, 35, 45, 237,
				87, 193, 81, 246, 239, 191, 10, 12, 138, 225, 29, 27, 160},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_K35OZTXL_WV7FH5SR_B3X2WJO4_UKAQSRNE_EQDJ5GQV_ZD43B3MP_WPPO6ANJ",
			AccountBalance: 0,
			NodePublicKey: []byte{204, 219, 12, 172, 185, 167, 187, 205, 196, 103, 169, 74, 242, 102, 223, 85, 5, 79, 233,
				110, 193, 163, 212, 73, 79, 64, 196, 240, 215, 155, 60, 105},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_GZDHPTO2_QMINSLSW_6SYWNR6M_S6BVZTII_3ODR5NTZ_365EGENK_23CFQTO6",
			AccountBalance: 0,
			NodePublicKey: []byte{72, 247, 106, 57, 38, 90, 173, 17, 89, 83, 206, 86, 148, 163, 211, 11, 183, 110, 92,
				16, 119, 58, 218, 209, 184, 85, 251, 132, 49, 30, 208, 157},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_64M4AZR5_3TI3ZDU6_EKNBJHCK_GFPTTFFP_TGKF43IG_M4HKF5NK_INHWKBY6",
			AccountBalance: 0,
			NodePublicKey: []byte{139, 153, 118, 92, 81, 76, 29, 214, 51, 174, 25, 7, 163, 64, 99, 88, 77, 132, 138,
				111, 52, 192, 75, 223, 230, 136, 216, 145, 120, 125, 154, 201},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
		{
			AccountAddress: "ZBC_A3QQPNAK_JML5Z74E_RSTUS3AN_RX2LAVUN_TNZ2Y76U_V654WYQ6_KGZK3SLN",
			AccountBalance: 0,
			NodePublicKey: []byte{76, 50, 6, 113, 42, 129, 225, 241, 121, 94, 75, 49, 129, 20, 197, 54, 241, 136, 168,
				242, 1, 11, 56, 63, 53, 129, 101, 239, 103, 13, 116, 34},
			LockedBalance:      0,
			ParticipationScore: 500000000000000000,
		},
	}
)
