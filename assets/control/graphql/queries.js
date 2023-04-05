import gql from 'graphql-tag';

//zarlagdsan tusul
export const GET_PROJECT_ZAR = gql`{
    tbl_project_zar{
        id
        garchig
        tailbar
        start_date
        end_date
        user_id
        attach_files
  }
}`;

export const GET_ZAR_BY_ID = gql`
query tbl_project_zar($user_id: String!){
  tbl_project_zar
  (
      filters: [
        { column: "user_id", condition: equals, value: $user_id }
      ]
  )
  {
         id
        garchig
        tailbar
        start_date
        end_date
        user_id
        attach_files
  }
}`;


//zarlagdsan tusuld nemex xvselt ywuulsan esex
export const GET_PROJECT_USER_CHECK = gql`
query z_project_user_check($user_id: String!) {
  z_project_user_check(
    filters: [{ column: "user_id", condition: equals, value: $user_id }]
  ) {
    user_id
    tailbar
    form_id
    project_status_id
    id
  }
}
`;


export const GET_PROJECT_ZAR_BY_ID = gql`
query tbl_project_zar($z_id: String!){
  tbl_project_zar
  (
      filters: [
        { column: "id", condition: equals, value: $z_id }
      ]
  )
  {
        id
        garchig
        tailbar
        start_date
        end_date
        user_id
        hayag
        tusliin_urgeljleh_hugtsaa
        sanhuujilt_baiguullaga
        sanhuujilt_huwi
        tailbar
        tusliin_urgeljleh_hugtsaa
        utas
        sub_project_shaardlaga{
          project_type
        }
        sub_project_type{
          project_type_id
        }
        sub_project_type_other{
          project_type

        }
  }
}`;
