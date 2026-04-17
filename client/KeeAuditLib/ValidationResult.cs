using KeeAuditLib.Specifications;
using KeePassLib;
using System;
using System.Collections.Generic;

namespace KeeAuditLib
{
    public class ValidationResult: IEquatable<ValidationResult>
    {
        private readonly PwEntry _entry;
        private readonly IList<Specification<PwEntry>> _failedSpecifications;

        public ValidationResult(PwEntry entry, IList<Specification<PwEntry>> failedSpecifications)
        {
            _entry = entry;
            _failedSpecifications = failedSpecifications;
        }

        public PwEntry Entry => _entry;
        public bool IsCompliant => _failedSpecifications.Count == 0;
        public IList<Specification<PwEntry>> FailedSpecifications => _failedSpecifications;

        public override bool Equals(object obj)
        {
            return Equals(obj as ValidationResult);
        }

        public bool Equals(ValidationResult other)
        {
            if (other == null) return false;
            return _entry.Uuid.Equals(other._entry.Uuid) && IsCompliant == other.IsCompliant;
        }

        public override int GetHashCode()
        {
            unchecked
            {
                int hash = 17;
                hash = hash * 23 + _entry.Uuid.GetHashCode();
                hash = hash * 23 + IsCompliant.GetHashCode();
                return hash;
            }
        }

        public static bool operator ==(ValidationResult left, ValidationResult right)
        {
            return Equals(left, right);
        }

        public static bool operator !=(ValidationResult left, ValidationResult right)
        {
            return !Equals(left, right);
        }
    }
}
