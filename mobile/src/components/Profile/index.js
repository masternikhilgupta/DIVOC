import React, {useEffect, useState} from "react";
import {BaseFormCard} from "../BaseFormCard";
import "./index.scss"
import {formatLoginDate} from "../../utils/date_utils";
import {appIndexDb} from "../../AppDatabase";
import {SyncFacade} from "../../SyncFacade";
import Button from "react-bootstrap/Button";
import {Messages} from "../../Base/Constants";
import {AuthSafeComponent} from "../../utils/keycloak";
import Col from "react-bootstrap/Col";
import Row from "react-bootstrap/Row";

function AuthSafeUserProfile({keycloak}) {
    const [userDetails, setUserDetails] = useState();

    useEffect(() => {
        appIndexDb.getUserDetails()
            .then((userDetails) => setUserDetails(userDetails))
            .catch((e) => {
            })

    }, [])


    return (
        <div className="profile-container">
            <BaseFormCard title="Profile">
                {userDetails && <div>
                    <div className="name mt-2">{userDetails.full_name}</div>
                    <div className="mt-2 mr-2">
                        <Row>
                            <Col>
                                <div className="subtitle label">Last logged in</div>
                                <div className="subtitle date">{formatLoginDate(userDetails.loginTime)}</div>
                            </Col>
                            <Col>
                                <div className="d-flex flex-column justify-content-end">
                                    <div className="subtitle label">Last sync</div>
                                    <div className="subtitle date">{SyncFacade.lastSyncedOn()}</div>
                                </div>
                            </Col>
                        </Row>
                    </div>
                    <hr/>
                    <div>
                        <div className="name">Facility Center</div>
                        <div className="mt-2">{userDetails.facilityDetails.facilityName}</div>
                        <div
                            className="subtitle mt-1">{userDetails.facilityDetails.address.addressLine1}{userDetails.facilityDetails.address.district ? "," : ""} {userDetails.facilityDetails.address.district}</div>
                        <div
                            className="subtitle">{userDetails.facilityDetails.address.state} {userDetails.facilityDetails.address.pincode}</div>
                    </div>
                    <div className="logout-container mt-5 d-flex justify-content-around">
                        <Button variant="outline-danger" onClick={() => {
                            if (navigator.onLine) {
                                SyncFacade
                                    .push()
                                    .catch((e) => console.log(e.message))
                                    .then(() => appIndexDb.clearEverything())
                                    .then((() => keycloak.logout()))
                                    .catch(e => {
                                        console.log(e.message)
                                        if (!navigator.onLine) {
                                            alert(Messages.NO_INTERNET_CONNECTION)
                                        }
                                    })
                            } else {
                                alert(Messages.NO_INTERNET_CONNECTION)
                            }
                        }}>Logout</Button>{" "}
                    </div>
                </div>}
            </BaseFormCard>
        </div>
    )
}

export function UserProfile(props) {
    return (
        <AuthSafeComponent>
            <AuthSafeUserProfile/>
        </AuthSafeComponent>
    );
}
